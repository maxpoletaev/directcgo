package codegen

import (
	"fmt"
	"go/types"
	"math/rand/v2"
)

var (
	amd64IntegerRegs = []string{"DI", "SI", "DX", "CX", "R8", "R9"}
	amd64FloatRegs   = []string{"X0", "X1", "X2", "X3", "X4", "X5", "X6", "X7"}
)

type amd64ArgClass uint8

const (
	amd64ArgClassNone amd64ArgClass = iota
	amd64ArgClassInteger
	amd64ArgClassSSE
	amd64ArgClassMemory
)

// AMD64 is an AMD64 code generator.
// Go assembly uses non-standard names for instructions and registers compared to Intel/AT&T syntax.
// See: https://golang.org/doc/asm and https://www.quasilyte.dev/blog/post/go-asm-complementary-reference/
type amd64 struct {
	ngpr int // next general purpose register number
	nsrn int // next simd and fp register number
	nsaa int // next stack argument address
}

func newAMD64() *amd64 {
	return &amd64{
		nsaa: 8,
	}
}

func (arch *amd64) resetState() {
	arch.ngpr = 0
	arch.nsrn = 0
	arch.nsaa = 8
}

func (arch *amd64) Name() string {
	return "amd64"
}

func (arch *amd64) totalArgsSize(fn *Function) (total int) {
	for _, arg := range fn.Args {
		size := typeSize(arg.Type)
		total = align(total, size)
		total += size
	}

	if fn.ReturnType != nil {
		size := typeSize(fn.ReturnType)
		total = align(total, size)
		total += size
	}

	return total
}

func (arch *amd64) loadInteger(buf *builder, arg *Argument, offset int, reg string) {
	size := typeSize(arg.Type)
	if isUnsigned(arg.Type) {
		switch size {
		case 8:
			buf.I("MOVQ", "%s+%d(FP), %s", arg.Name, offset, reg)
		case 4:
			buf.I("MOVLQZX", "%s+%d(FP), %s", arg.Name, offset, reg)
		case 2:
			buf.I("MOVWQZX", "%s+%d(FP), %s", arg.Name, offset, reg)
		case 1:
			buf.I("MOVBQZX", "%s+%d(FP), %s", arg.Name, offset, reg)
		default:
			panic(fmt.Sprintf("unknown int size: %d", size))
		}
	} else {
		switch size {
		case 8:
			buf.I("MOVQ", "%s+%d(FP), %s", arg.Name, offset, reg)
		case 4:
			buf.I("MOVLQSX", "%s+%d(FP), %s", arg.Name, offset, reg)
		case 2:
			buf.I("MOVWQSX", "%s+%d(FP), %s", arg.Name, offset, reg)
		case 1:
			buf.I("MOVBQSX", "%s+%d(FP), %s", arg.Name, offset, reg)
		default:
			panic(fmt.Sprintf("unknown int size: %d", size))
		}
	}
}

func (arch *amd64) loadFloat(buf *builder, arg *Argument, offset int, reg string) {
	size := typeSize(arg.Type)

	switch size {
	case 4:
		buf.I("MOVSS", "%s+%d(FP), %s", arg.Name, offset, reg)
	case 8:
		buf.I("MOVSD", "%s+%d(FP), %s", arg.Name, offset, reg)
	default:
		panic(fmt.Sprintf("unknown float size: %d", size))
	}
}

func (arch *amd64) mergeClasses(c1, c2 amd64ArgClass) amd64ArgClass {
	switch {
	case c1 == c2:
		return c1
	case c1 == amd64ArgClassInteger,
		c2 == amd64ArgClassInteger:
		return amd64ArgClassInteger
	case c1 == amd64ArgClassMemory,
		c2 == amd64ArgClassMemory:
		return amd64ArgClassMemory
	default:
		return amd64ArgClassSSE
	}
}

func (arch *amd64) classifyType(ty types.Type) amd64ArgClass {
	switch {
	case isFloatingPoint(ty):
		return amd64ArgClassSSE
	case isInteger(ty) && typeSize(ty) <= 8:
		return amd64ArgClassInteger
	case isComposite(ty):
		fields := getFields(ty)
		class := amd64ArgClassNone

		for _, field := range fields {
			fieldClass := arch.classifyType(field)
			class = arch.mergeClasses(class, fieldClass)
		}

		return class
	}

	panic(fmt.Sprintf("unhandled type: %s", ty))
}

func (arch *amd64) loadSmallStruct(buf *builder, arg *Argument, offset int) int {
	size := typeSize(arg.Type)
	fields := getFields(arg.Type)
	nEightbytes := (size + 7) / 8

	classes := make([]amd64ArgClass, nEightbytes)
	fieldOffset := 0

	for _, field := range fields {
		fieldSize := typeSize(field)
		fieldOffset = align(fieldOffset, fieldSize)
		eightbyte := fieldOffset / 8

		class := arch.classifyType(field)
		classes[eightbyte] = arch.mergeClasses(classes[eightbyte], class)

		fieldOffset += fieldSize
	}

	for i := 0; i < nEightbytes; i++ {
		switch classes[i] {
		case amd64ArgClassSSE:
			reg := amd64FloatRegs[arch.nsrn]
			buf.I("MOVQ", "%s+%d(FP), %s", arg.Name, offset, reg)
			arch.nsrn++
			offset += 8
		case amd64ArgClassInteger:
			reg := amd64IntegerRegs[arch.ngpr]
			buf.I("MOVQ", "%s+%d(FP), %s", arg.Name, offset, reg)
			arch.ngpr++
			offset += 8
		default:
			panic(fmt.Sprintf("unhandled eightbyte class: %d", classes[i]))
		}
	}

	return offset
}

func (arch *amd64) loadArg(buf *builder, arg *Argument, offset int) int {
	size := typeSize(arg.Type)

	if isInteger(arg.Type) && arch.ngpr < len(amd64IntegerRegs) {
		reg := amd64IntegerRegs[arch.ngpr]
		offset = align(offset, size)
		arch.loadInteger(buf, arg, offset, reg)
		arch.ngpr++
		return offset + size
	}

	if isFloatingPoint(arg.Type) && arch.nsrn < len(amd64FloatRegs) {
		reg := amd64FloatRegs[arch.nsrn]
		offset = align(offset, size)
		arch.loadFloat(buf, arg, offset, reg)
		arch.nsrn++
		return offset + size
	}

	// Small composite types are divided into eightbytes (8-byte chunks).
	// Each eightbyte can be classified as integer, sse, or memory.
	if isComposite(arg.Type) && size <= 16 {
		return arch.loadSmallStruct(buf, arg, offset)
	}

	panic(fmt.Sprintf("unhandled argument: %s", arg.Name))
}

func (arch *amd64) storeReturn(buf *builder, arg *Argument, offset int) int {
	size := typeSize(arg.Type)
	offset = align(offset, size)

	switch {
	case isInteger(arg.Type):
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
	case isFloatingPoint(arg.Type):
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
		panic(fmt.Sprintf("unsupported return type: %T", arg.Type))
	}

	offset += size
	return offset
}

func (arch *amd64) GenerateFunc(buf *builder, f *Function) {
	arch.resetState()

	buf.S("// %s", f.Signature)
	buf.S("TEXT ·%s(SB), $%d-%d", f.Name, defaultFrameSize, arch.totalArgsSize(f))
	buf.I("MOVQ", "%s+0(FP), AX", f.Args[0].Name)

	// Load arguments
	offset := 8
	for i := 1; i < len(f.Args); i++ {
		offset = arch.loadArg(buf, &f.Args[i], offset)
	}

	seed := [32]byte{}
	copy(seed[:], f.Name)
	rnd := rand.New(rand.NewChaCha8(seed))

	// Set frame guard
	guardValue := rnd.Uint32()
	buf.I("MOVL", "$0x%X, R10", guardValue)
	buf.I("MOVL", "R10, 8(SP)")

	// Stack adjustment
	buf.I("MOVQ", "SP, R12")
	buf.I("LEAQ", "%d(SP), SP", defaultFrameSize)
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
		arch.storeReturn(buf, &Argument{
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
