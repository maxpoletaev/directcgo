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

func isTypeUnsigned(t types.Type) bool {
	switch t := t.Underlying().(type) {
	case *types.Basic:
		return t.Info()&types.IsUnsigned != 0
	default:
		return true
	}
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

func isStructHFA(t types.Type) bool {
	st := t.Underlying().(*types.Struct)
	for i := 0; i < st.NumFields(); i++ {
		if getArgKind(st.Field(i).Type()) != ArgFloat {
			return false
		}
	}
	return true
}
