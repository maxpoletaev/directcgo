package codegen

import (
	"fmt"
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

func (arch *AMD64) totalArgsSize(fn *Function) (total int) {
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

func (arch *AMD64) loadInteger(buf *builder, arg *Argument, offset int, reg string) {
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

func (arch *AMD64) loadFloat(buf *builder, arg *Argument, offset int, reg string) {
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

func (arch *AMD64) loadHFA(buf *builder, arg *Argument, offset int, regs []string) {
	fields := getFields(arg.Type)

	for i, field := range fields {
		reg := regs[i]
		size := typeSize(field)
		offset = align(offset, size)

		arch.loadFloat(buf, &Argument{
			Name: arg.Name + "_" + strconv.Itoa(i),
			Type: field,
		}, offset, reg)

		offset += size
	}
}

func (arch *AMD64) loadMultiReg(buf *builder, arg *Argument, offset int, regs []string) {
	for _, reg := range regs {
		buf.I("MOVQ", "%s+%d(FP), %s", arg.Name, offset, reg)
		offset += 8
	}
}

func (arch *AMD64) loadArg(buf *builder, arg *Argument, offset int) int {
	intRegs := [6]string{"DI", "SI", "DX", "CX", "R8", "R9"}
	size := typeSize(arg.Type)

	if isInteger(arg.Type) && arch.intCount < AMD64IntRegs {
		reg := intRegs[arch.intCount]
		offset = align(offset, size)
		arch.loadInteger(buf, arg, offset, reg)
		arch.intCount++
		return offset + size
	}

	if isFloatingPoint(arg.Type) && arch.floatCount < AMD64FloatRegs {
		reg := fmt.Sprintf("X%d", arch.floatCount)
		offset = align(offset, size)
		arch.loadFloat(buf, arg, offset, reg)
		arch.floatCount++
		return offset + size
	}

	if hfa, _ := isHFA(arg.Type); hfa {
		numFields := getFieldCount(arg.Type)

		if arch.floatCount+numFields <= AMD64FloatRegs {
			regs := make([]string, numFields)
			for i := range regs {
				regs[i] = fmt.Sprintf("X%d", arch.floatCount+i)
			}

			arch.loadHFA(buf, arg, offset, regs)
			arch.floatCount += numFields
			return offset + size
		}
	}

	if isComposite(arg.Type) {
		if size/8 <= AMD64IntRegs-arch.intCount {
			regs := intRegs[arch.intCount : arch.intCount+size/8]
			arch.loadMultiReg(buf, arg, offset, regs[:size/8])
			arch.intCount += size / 8
			return offset + size
		}
	}

	panic(fmt.Sprintf("unhandled argument: %s", arg.Name))
}

func (arch *AMD64) storeReturn(buf *builder, arg *Argument, offset int) int {
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

func (arch *AMD64) GenerateFunc(buf *builder, f *Function) {
	arch.resetState()

	buf.S("// %s", f.Signature)
	buf.S("TEXT ·%s(SB), $%d-%d", f.Name, AMD64FrameSize, arch.totalArgsSize(f))
	buf.I("MOVQ", "%s+0(FP), AX", f.Args[0].Name)

	// Load arguments
	offset := 8
	for i := 1; i < len(f.Args); i++ {
		offset = arch.loadArg(buf, &f.Args[i], offset)
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
