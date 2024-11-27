package codegen

import (
	"fmt"
	"go/types"
	"math/rand/v2"
)

const (
	ARM64FrameSize = 65536
	ARM64IntRegs   = 8 // R0-R7
	ARM64FloatRegs = 8 // F0-F7
)

// ARM64 is an ARM64 code generator.
// Go assembly uses non-standard names for instructions and registers,
// e.g., ldr is called MOVD, etc. See: https://pkg.go.dev/cmd/internal/obj/arm64
// Procedure call: https://developer.arm.com/documentation/102374/0102/Procedure-Call-Standard
type ARM64 struct {
	intCount   int
	floatCount int
}

func NewArm64() *ARM64 {
	return &ARM64{}
}

func (arch *ARM64) resetState() {
	arch.intCount = 0
	arch.floatCount = 0
}

func (arch *ARM64) Name() string {
	return "arm64"
}

func (arch *ARM64) typeSize(t types.Type) int {
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

func (arch *ARM64) totalArgsSize(fn *Function) (total int) {
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

func (arch *ARM64) nextReg(kind ArgKind) string {
	switch kind {
	case ArgInt:
		if arch.intCount >= ARM64IntRegs {
			panic("out of integer registers")
		}
		reg := fmt.Sprintf("R%d", arch.intCount)
		arch.intCount++
		return reg
	case ArgFloat:
		if arch.floatCount >= ARM64FloatRegs {
			panic("out of float registers")
		}
		reg := fmt.Sprintf("F%d", arch.floatCount)
		arch.floatCount++
		return reg
	default:
		panic("unknown argument kind")
	}
}

func (arch *ARM64) loadIntArg(buf *builder, unsigned bool, size int, name string, offset int, reg string) {
	if unsigned {
		switch size {
		case 8:
			buf.I("MOVD", "%s+%d(FP), %s", name, offset, reg)
		case 4:
			buf.I("MOVWU", "%s+%d(FP), %s", name, offset, reg)
		case 2:
			buf.I("MOVHU", "%s+%d(FP), %s", name, offset, reg)
		case 1:
			buf.I("MOVBU", "%s+%d(FP), %s", name, offset, reg)
		default:
			panic(fmt.Sprintf("unknown int size: %d", size))
		}
	} else {
		switch size {
		case 8:
			buf.I("MOVD", "%s+%d(FP), %s", name, offset, reg)
		case 4:
			buf.I("MOVW", "%s+%d(FP), %s", name, offset, reg)
		case 2:
			buf.I("MOVH", "%s+%d(FP), %s", name, offset, reg)
		case 1:
			buf.I("MOVB", "%s+%d(FP), %s", name, offset, reg)
		default:
			panic(fmt.Sprintf("unknown int size: %d", size))
		}
	}
}

func (arch *ARM64) loadFloatArg(buf *builder, size int, name string, offset int, reg string) {
	switch size {
	case 4:
		buf.I("FMOVS", "%s+%d(FP), %s", name, offset, reg)
	case 8:
		buf.I("FMOVD", "%s+%d(FP), %s", name, offset, reg)
	default:
		panic(fmt.Sprintf("unknown float size: %d", size))
	}
}

func (arch *ARM64) argLoad(buf *builder, arg *Argument, offset int) int {
	size := arch.typeSize(arg.Type)
	kind := getArgKind(arg.Type)

	switch kind {
	case ArgInt:
		reg := arch.nextReg(kind)
		offset = align(offset, size)
		unsigned := isTypeUnsigned(arg.Type)
		arch.loadIntArg(buf, unsigned, size, arg.Name, offset, reg)
	case ArgFloat:
		reg := arch.nextReg(kind)
		offset = align(offset, size)
		arch.loadFloatArg(buf, size, arg.Name, offset, reg)
	case ArgStruct:
		if size <= 16 {
			arch.loadSmallStruct(buf, arg, offset)
		} else {
			panic("struct size > 16 not supported")
		}
	default:
		panic(fmt.Sprintf("unknown argument kind: %d", kind))
	}

	offset += size
	return offset
}

func (arch *ARM64) loadSmallStruct(buf *builder, arg *Argument, offset int) {
	st := arg.Type.Underlying().(*types.Struct)
	structSize := arch.typeSize(st)

	var (
		rem        = structSize
		currOffset = offset
		chunkSize  = 8
	)

	for rem > 0 {
		reg := arch.nextReg(ArgInt)
		arch.loadIntArg(buf, false, chunkSize, arg.Name, currOffset, reg)
		currOffset += chunkSize
		rem -= chunkSize
	}
}

func (arch *ARM64) retStore(buf *builder, arg *Argument, offset int) int {
	size := arch.typeSize(arg.Type)
	kind := getArgKind(arg.Type)
	offset = align(offset, size)

	switch kind {
	case ArgInt:
		reg := "R0"
		switch size {
		case 1:
			buf.I("MOVB", "%s, %s+%d(FP)", reg, arg.Name, offset)
		case 2:
			buf.I("MOVH", "%s, %s+%d(FP)", reg, arg.Name, offset)
		case 4:
			buf.I("MOVW", "%s, %s+%d(FP)", reg, arg.Name, offset)
		case 8:
			buf.I("MOVD", "%s, %s+%d(FP)", reg, arg.Name, offset)
		default:
			panic(fmt.Sprintf("unknown int size: %d", size))
		}
	case ArgFloat:
		reg := "F0"
		switch size {
		case 4:
			buf.I("FMOVS", "%s, %s+%d(FP)", reg, arg.Name, offset)
		case 8:
			buf.I("FMOVD", "%s, %s+%d(FP)", reg, arg.Name, offset)
		default:
			panic(fmt.Sprintf("unknown float size: %d", size))
		}
	default:
		panic(fmt.Sprintf("unknown argument kind: %d", kind))
	}

	offset += size
	return offset
}

func (arch *ARM64) GenerateFunc(buf *builder, f *Function) {
	arch.resetState()

	buf.S("// %s", f.Signature)
	buf.S("TEXT ·%s(SB), $%d-%d", f.Name, ARM64FrameSize, arch.totalArgsSize(f))
	buf.I("MOVD", "%s+0(FP), R9", f.Args[0].Name)

	// Load arguments
	offset := 8
	for i := 1; i < len(f.Args); i++ {
		offset = arch.argLoad(buf, &f.Args[i], offset)
	}

	// Set frame guard
	guardValue := rand.Uint32()
	buf.I("MOVD", "$0x%X, R10", guardValue)
	buf.I("MOVD", "R10, 8(RSP)")

	// Stack adjustment
	buf.I("MOVD", "RSP, R20")
	buf.I("MOVD", "$%d, R10", ARM64FrameSize)
	buf.I("ADD", "R10, RSP")
	buf.I("MOVD", "RSP, R10")
	buf.I("AND", "$~15, R10, RSP")

	// Call function
	buf.I("BL", "(R9)")
	buf.I("MOVD", "R20, RSP")

	// Check frame guard
	buf.I("MOVD", "8(RSP), R10")
	buf.I("MOVD", "$0x%X, R11", guardValue)
	buf.I("CMP", "R10, R11")
	buf.I("BNE", "overflow")

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
