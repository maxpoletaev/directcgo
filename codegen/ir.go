package codegen

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"go/types"
	"log"

	"golang.org/x/tools/go/packages"
)

type ArgKind uint8

const (
	ArgInt ArgKind = iota
	ArgFloat
	ArgStruct
)

type Argument struct {
	Name string
	Type types.Type
}

type Function struct {
	Name       string
	Args       []Argument
	ReturnType types.Type
	Signature  string
}

func parsePackage(pkg *packages.Package) ([]*Function, error) {
	var funcs []*Function
	var parseErr error

	for _, file := range pkg.Syntax {
		ast.Inspect(file, func(n ast.Node) bool {
			if parseErr != nil {
				return false
			}

			if decl, ok := n.(*ast.FuncDecl); ok {
				if decl.Body != nil || decl.Name.Name[0] == '_' || !hasNoEscapeComment(decl) {
					return false
				}

				fn, err := parseFunction(decl, pkg.TypesInfo)
				if err != nil {
					parseErr = fmt.Errorf("failed to process function %s: %v", decl.Name.Name, err)
					return false
				}

				log.Printf("found %s", fn.Signature)
				funcs = append(funcs, fn)
			}

			return true
		})
	}

	if parseErr != nil {
		return nil, parseErr
	}

	return funcs, nil
}

func hasNoEscapeComment(decl *ast.FuncDecl) bool {
	if decl.Doc == nil {
		return false
	}
	for _, comment := range decl.Doc.List {
		if comment.Text == "//go:noescape" {
			return true
		}
	}
	return false
}

func parseFunction(decl *ast.FuncDecl, info *types.Info) (*Function, error) {
	f := &Function{
		Signature: getFuncSignature(decl),
		Name:      decl.Name.Name,
	}

	// Collect arguments
	if decl.Type.Params != nil {
		for _, field := range decl.Type.Params.List {
			typ := info.TypeOf(field.Type)

			for _, name := range field.Names {
				f.Args = append(f.Args, Argument{
					Name: name.Name,
					Type: typ,
				})
			}
		}
	}

	// Collect return type
	if decl.Type.Results != nil {
		if len(decl.Type.Results.List) > 1 {
			return nil, fmt.Errorf("multiple return values")
		}

		for _, field := range decl.Type.Results.List {
			typ := info.TypeOf(field.Type)
			f.ReturnType = typ
		}
	}

	return f, nil
}

func getFuncSignature(f *ast.FuncDecl) string {
	var buf bytes.Buffer

	err := printer.Fprint(&buf, token.NewFileSet(), f.Type)
	if err != nil {
		panic(err)
	}

	return f.Name.Name + " " + buf.String()
}

func getArgKind(t types.Type) ArgKind {
	switch t := t.Underlying().(type) {
	case *types.Basic:
		switch t.Kind() {
		case types.Float32, types.Float64:
			return ArgFloat
		default:
			return ArgInt
		}
	case *types.Pointer:
		return ArgInt
	case *types.Struct:
		return ArgStruct
	default:
		panic(fmt.Sprintf("unsupported type: %T", t))
	}
}

func typeSize(t types.Type) int {
	switch t := t.Underlying().(type) {
	case *types.Basic:
		switch t.Kind() {
		case types.Bool, types.Int8, types.Uint8:
			return 1
		case types.Int16, types.Uint16:
			return 2
		case types.Int32, types.Uint32, types.Float32:
			return 4
		case types.Int64, types.Uint64, types.Float64:
			return 8
		case types.UnsafePointer:
			return 8
		default:
			panic(fmt.Sprintf("unknown basic type: %s", t))
		}
	case *types.Pointer:
		return 8
	case *types.Struct:
		var size int
		for i := 0; i < t.NumFields(); i++ {
			s := typeSize(t.Field(i).Type())
			size = align(size, s)
			size += s
		}
		return size
	case *types.Array:
		return typeSize(t.Elem()) * int(t.Len())
	}
	panic(fmt.Sprintf("unsupported type: %T", t))
}

func isShortVector(t types.Type) bool {
	return false
}

func isComposite(t types.Type) bool {
	switch t.Underlying().(type) {
	case *types.Struct, *types.Array:
		return true
	default:
		return false
	}
}

func isInteger(t types.Type) bool {
	if basic, ok := t.Underlying().(*types.Basic); ok {
		return basic.Info()&types.IsInteger != 0
	}
	return false
}

func isUnsigned(t types.Type) bool {
	if basic, ok := t.Underlying().(*types.Basic); ok {
		return basic.Info()&types.IsUnsigned != 0
	}
	return false
}

func isFloatingPoint(t types.Type) bool {
	if basic, ok := t.Underlying().(*types.Basic); ok {
		return basic.Info()&types.IsFloat != 0
	}
	return false
}

func isPointer(t types.Type) bool {
	switch t := t.Underlying().(type) {
	case *types.Basic:
		if t.Kind() == types.UnsafePointer {
			return true
		}
	case *types.Pointer:
		return true
	}
	return false
}

func getFieldCount(t types.Type) int {
	switch t := t.Underlying().(type) {
	case *types.Struct:
		return t.NumFields()
	case *types.Array:
		return int(t.Len())
	default:
		panic(fmt.Sprintf("not a composite type: %T", t))
	}
}

func getFields(t types.Type) []types.Type {
	switch typ := t.Underlying().(type) {
	case *types.Struct:
		fields := make([]types.Type, typ.NumFields())
		for i := 0; i < typ.NumFields(); i++ {
			fields[i] = typ.Field(i).Type()
		}
		return fields
	case *types.Array:
		// For arrays, return a slice with the element type repeated length times
		fields := make([]types.Type, typ.Len())
		for i := range fields {
			fields[i] = typ.Elem()
		}
		return fields
	default:
		panic(fmt.Sprintf("not a composite type: %T", t))
	}
}

func isHFA(t types.Type) (_, sameType bool) {
	if !isComposite(t) {
		return false, false
	}

	fields := getFields(t)
	if len(fields) == 0 || len(fields) > 4 {
		return false, false
	}

	firstField := fields[0]
	sameType = true

	for _, field := range fields {
		if !isFloatingPoint(field) {
			return false, false
		}

		if field != firstField {
			sameType = false
		}
	}

	return true, sameType
}

func isHVA(t types.Type) bool {
	return false
}
