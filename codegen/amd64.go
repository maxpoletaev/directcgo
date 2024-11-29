package codegen

import (
	"fmt"
	"go/types"
	"math/rand/v2"
	"strconv"
)

const (
	AMD64FrameSize = 65536
	AMD64IntRegs   = 6 // DI, SI, DX, CX, R8, R9
	AMD64FloatRegs = 8 // XMM0-XMM7
)

// AMD64 is an AMD64 code generator.
// Go assembly uses non-standard names for instructions and registers compared to Intel/AT&T syntax.
// See: https://golang.org/doc/asm and https://www.quasilyte.dev/blog/post/go-asm-complementary-reference/
type AMD64 struct {
	intCount   int
	floatCount int
}

func NewAmd64() *AMD64 {
	return &AMD64{}
}

func (arch *AMD64) resetState() {
	arch.intCount = 0
	arch.floatCount = 0
}

func (arch *AMD64) Name() string {
	return "amd64"
}

func (arch *AMD64) typeSize(t types.Type) int {
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
			s := arch.typeSize(t.Field(i).Type())
			size = align(size, s)
			size += s
		}
		return size
	}
	panic(fmt.Sprintf("unsupported type: %T", t))
}

func (arch *AMD64) totalArgsSize(fn *Function) (total int) {
	for _, arg := range fn.Args {
		size := arch.typeSize(arg.Type)
		total = align(total, size)
		total += size
	}

	if fn.ReturnType != nil {
		size := arch.typeSize(fn.ReturnType)
		total = align(total, size)
		total += size
	}

	return total
}

func (arch *AMD64) nextReg(kind ArgKind) string {
	intRegs := [6]string{"DI", "SI", "DX", "CX", "R8", "R9"}

	switch kind {
	case ArgInt:
		if arch.intCount >= len(intRegs) {
			panic("out of integer registers")
		}
		reg := intRegs[arch.intCount]
		arch.intCount++
		return reg
	case ArgFloat:
		if arch.floatCount >= AMD64FloatRegs {
			panic("out of float registers")
		}
		reg := fmt.Sprintf("X%d", arch.floatCount)
		arch.floatCount++
		return reg
	default:
		panic(fmt.Sprintf("unknown argument kind: %d", kind))
	}
}

func (arch *AMD64) loadIntArg(buf *builder, unsigned bool, size int, name string, offset int, reg string) {
	if unsigned {
		switch size {
		case 8:
			buf.I("MOVQ", "%s+%d(FP), %s", name, offset, reg)
		case 4:
			buf.I("MOVLQZX", "%s+%d(FP), %s", name, offset, reg)
		case 2:
			buf.I("MOVWQZX", "%s+%d(FP), %s", name, offset, reg)
		case 1:
			buf.I("MOVBQZX", "%s+%d(FP), %s", name, offset, reg)
		default:
			panic(fmt.Sprintf("unknown int size: %d", size))
		}
	} else {
		switch size {
		case 8:
			buf.I("MOVQ", "%s+%d(FP), %s", name, offset, reg)
		case 4:
			buf.I("MOVLQSX", "%s+%d(FP), %s", name, offset, reg)
		case 2:
			buf.I("MOVWQSX", "%s+%d(FP), %s", name, offset, reg)
		case 1:
			buf.I("MOVBQSX", "%s+%d(FP), %s", name, offset, reg)
		default:
			panic(fmt.Sprintf("unknown int size: %d", size))
		}
	}
}

func (arch *AMD64) loadFloatArg(buf *builder, size int, name string, offset int, reg string) {
	switch size {
	case 4:
		buf.I("MOVSS", "%s+%d(FP), %s", name, offset, reg)
	case 8:
		buf.I("MOVSD", "%s+%d(FP), %s", name, offset, reg)
	default:
		panic(fmt.Sprintf("unknown float size: %d", size))
	}
}

func (arch *AMD64) loadSmallStructArg(buf *builder, arg *Argument, offset int) {
	st := arg.Type.Underlying().(*types.Struct)
	structSize := arch.typeSize(st)

	if hfa, _ := isHFA(arg.Type); hfa {
		localOffset := offset
		for i := 0; i < st.NumFields(); i++ {
			field := st.Field(i)
			size := arch.typeSize(field.Type())
			localOffset = align(localOffset, size)
			name := arg.Name + "_" + field.Name()
			reg := arch.nextReg(ArgFloat)
			arch.loadFloatArg(buf, size, name, localOffset, reg)
			localOffset += size
		}
	} else {
		var (
			rem         = structSize
			localOffset = offset
			chunkSize   = 8
		)

		i := 0
		for rem > 0 {
			reg := arch.nextReg(ArgInt)
			argName := arg.Name + "_" + strconv.Itoa(i)
			buf.I("MOVQ", "%s+%d(FP), %s", argName, localOffset, reg)
			localOffset += chunkSize
			rem -= chunkSize
			i++
		}
	}
}

func (arch *AMD64) argLoad(buf *builder, arg *Argument, offset int) int {
	size := arch.typeSize(arg.Type)
	kind := getArgKind(arg.Type)

	switch kind {
	case ArgInt:
		reg := arch.nextReg(kind)
		offset = align(offset, size)
		unsigned := isUnsigned(arg.Type)
		arch.loadIntArg(buf, unsigned, size, arg.Name, offset, reg)
	case ArgFloat:
		reg := arch.nextReg(kind)
		offset = align(offset, size)
		arch.loadFloatArg(buf, size, arg.Name, offset, reg)
	case ArgStruct:
		if size <= 16 {
			arch.loadSmallStructArg(buf, arg, offset)
		} else {
			reg := arch.nextReg(ArgInt)
			buf.I("LEAQ", "%s+%d(FP), %s", arg.Name, offset, reg)
		}
	default:
		panic(fmt.Sprintf("unknown argument kind: %d", kind))
	}

	offset += size
	return offset
}

func (arch *AMD64) retStore(buf *builder, arg *Argument, offset int) int {
	size := arch.typeSize(arg.Type)
	kind := getArgKind(arg.Type)
	offset = align(offset, size)

	switch kind {
	case ArgInt:
		reg := "AX"
		switch size {
		case 1:
			buf.I("MOVB", "%s, %s+%d(FP)", reg, arg.Name, offset)
		case 2:
			buf.I("MOVW", "%s, %s+%d(FP)", reg, arg.Name, offset)
		case 4:
			buf.I("MOVL", "%s, %s+%d(FP)", reg, arg.Name, offset)
		case 8:
			buf.I("MOVQ", "%s, %s+%d(FP)", reg, arg.Name, offset)
		default:
			panic(fmt.Sprintf("unknown int size: %d", size))
		}
	case ArgFloat:
		reg := "X0"
		switch size {
		case 4:
			buf.I("MOVSS", "%s, %s+%d(FP)", reg, arg.Name, offset)
		case 8:
			buf.I("MOVSD", "%s, %s+%d(FP)", reg, arg.Name, offset)
		default:
			panic(fmt.Sprintf("unknown float size: %d", size))
		}
	default:
		panic(fmt.Sprintf("unknown argument kind: %d", kind))
	}

	offset += size
	return offset
}

func (arch *AMD64) GenerateFunc(buf *builder, f *Function) {
	arch.resetState()

	buf.S("// %s", f.Signature)
	buf.S("TEXT ·%s(SB), $%d-%d", f.Name, AMD64FrameSize, arch.totalArgsSize(f))
	buf.I("MOVQ", "%s+0(FP), AX", f.Args[0].Name)

	// Load arguments
	offset := 8
	for i := 1; i < len(f.Args); i++ {
		offset = arch.argLoad(buf, &f.Args[i], offset)
	}

	// Set frame guard
	guardValue := rand.Uint32()
	buf.I("MOVL", "$0x%X, R10", guardValue)
	buf.I("MOVL", "R10, 8(SP)")

	// Stack adjustment
	buf.I("MOVQ", "SP, R12")
	buf.I("LEAQ", "%d(SP), SP", AMD64FrameSize)
	buf.I("ANDQ", "$~15, SP")

	// Call function
	buf.I("CALL", "AX")
	buf.I("MOVQ", "R12, SP")

	// Check frame guard
	buf.I("MOVL", "8(SP), R10")
	buf.I("CMPL", "R10, $0x%X", guardValue)
	buf.I("JNE", "overflow")

	// Store return value
	if f.ReturnType != nil {
		arch.retStore(buf, &Argument{
			Type: f.ReturnType,
			Name: "ret",
		}, offset)
	}

	buf.I("RET", "")

	// Overflow handler
	buf.S("overflow:")
	buf.I("CALL", "runtime·abort(SB)")
	buf.I("RET", "")
}
