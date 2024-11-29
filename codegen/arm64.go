package codegen

import (
	"fmt"
	"math/rand/v2"
)

const (
	ARM64FrameSize = 65536
	ARM64IntRegs   = 8 // R0-R7
	ARM64FloatRegs = 8 // F0-F7
)

type allocResult struct {
	regs        []string
	isFloat     bool
	onStack     bool
	stackOffset int
	size        int
}

// ARM64 is an ARM64 code generator.
// Go assembly uses non-standard names for instructions and registers,
// e.g., ldr is called MOVD, etc. See: https://pkg.go.dev/cmd/internal/obj/arm64
// Procedure call: https://developer.arm.com/documentation/102374/0102/Procedure-Call-Standard
type ARM64 struct {
	ngrn int // Next General-purpose Register Number
	nsrn int // Next SIMD and FP Register Number
	nsaa int // Next Stack Argument Address
}

func NewArm64() *ARM64 {
	return &ARM64{
		nsaa: 8,
	}
}

func (arch *ARM64) resetState() {
	arch.ngrn = 0
	arch.nsrn = 0
	arch.nsaa = 8
}

func (arch *ARM64) Name() string {
	return "arm64"
}

func (arch *ARM64) totalArgsSize(fn *Function) (total int) {
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

func (arch *ARM64) loadInteger(buf *builder, arg *Argument, offset int, reg string) {
	size := typeSize(arg.Type)
	if isUnsigned(arg.Type) {
		switch size {
		case 1:
			buf.I("MOVBU", "%s+%d(FP), %s", arg.Name, offset, reg)
		case 2:
			buf.I("MOVHU", "%s+%d(FP), %s", arg.Name, offset, reg)
		case 4:
			buf.I("MOVWU", "%s+%d(FP), %s", arg.Name, offset, reg)
		case 8:
			buf.I("MOVD", "%s+%d(FP), %s", arg.Name, offset, reg)
		default:
			panic(fmt.Sprintf("unknown int size: %d", size))
		}
	} else {
		switch size {
		case 1:
			buf.I("MOVB", "%s+%d(FP), %s", arg.Name, offset, reg)
		case 2:
			buf.I("MOVH", "%s+%d(FP), %s", arg.Name, offset, reg)
		case 4:
			buf.I("MOVW", "%s+%d(FP), %s", arg.Name, offset, reg)
		case 8:
			buf.I("MOVD", "%s+%d(FP), %s", arg.Name, offset, reg)
		default:
			panic(fmt.Sprintf("unknown int size: %d", size))
		}
	}
}

func (arch *ARM64) loadFloat(buf *builder, arg *Argument, offset int, reg string) {
	size := typeSize(arg.Type)
	switch size {
	case 4:
		buf.I("FMOVS", "%s+%d(FP), %s", arg.Name, offset, reg)
	case 8:
		buf.I("FMOVD", "%s+%d(FP), %s", arg.Name, offset, reg)
	default:
		panic(fmt.Sprintf("unknown float size: %d", size))
	}
}

func (arch *ARM64) loadMultiReg(buf *builder, arg *Argument, offset int, regs []string) {
	for _, reg := range regs {
		buf.I("MOVD", "%s+%d(FP), %s", arg.Name, offset, reg)
		offset += 8
	}
}

func (arch *ARM64) loadHFA(buf *builder, arg *Argument, offset int, regs []string) {
	fields := getFields(arg.Type)

	for i, field := range fields {
		fieldSize := typeSize(field)
		offset = align(offset, fieldSize)

		arch.loadFloat(buf, &Argument{
			Name: fmt.Sprintf("%s_%d", arg.Name, i),
			Type: field,
		}, offset, regs[i])

		offset += fieldSize
	}
}

func (arch *ARM64) loadArg(buf *builder, arg *Argument, offset int) int {
	ty := arg.Type
	size := typeSize(ty)

	if isInteger(ty) && arch.ngrn < ARM64IntRegs {
		offset = align(offset, size)
		reg := fmt.Sprintf("R%d", arch.ngrn)
		arch.ngrn++

		arch.loadInteger(buf, arg, offset, reg)
		return offset + size
	}

	if isFloatingPoint(ty) && arch.nsrn < ARM64FloatRegs {
		offset = align(offset, size)
		reg := fmt.Sprintf("F%d", arch.nsrn)
		arch.nsrn++

		arch.loadFloat(buf, arg, offset, reg)
		return offset + size
	}

	if isHFA(ty) {
		numFields := getFieldCount(ty)

		if arch.nsrn+numFields <= ARM64FloatRegs {
			regs := make([]string, numFields)
			for i := 0; i < len(regs); i++ {
				regs[i] = fmt.Sprintf("F%d", arch.nsrn)
				arch.nsrn++
			}

			arch.loadHFA(buf, arg, offset, regs)
			return offset + size
		}
	}

	if isComposite(ty) {
		if size/8 <= ARM64IntRegs-arch.ngrn {
			regs := make([]string, (size+7)/8)
			for i := 0; i < len(regs); i++ {
				regs[i] = fmt.Sprintf("R%d", arch.ngrn)
				arch.ngrn++
			}

			arch.loadMultiReg(buf, arg, offset, regs)
			return offset + size
		}
	}

	panic(fmt.Sprintf("unhandled argument: %s", arg.Name))
}

func (arch *ARM64) storeReturn(buf *builder, arg *Argument, offset int) int {
	size := typeSize(arg.Type)
	offset = align(offset, size)

	switch {
	case isFloatingPoint(arg.Type):
		reg := "F0"
		switch size {
		case 4:
			buf.I("FMOVS", "%s, %s+%d(FP)", reg, arg.Name, offset)
		case 8:
			buf.I("FMOVD", "%s, %s+%d(FP)", reg, arg.Name, offset)
		default:
			panic(fmt.Sprintf("unknown float size: %d", size))
		}
	case isInteger(arg.Type):
		reg := "R0"
		if isUnsigned(arg.Type) {
			switch size {
			case 1:
				buf.I("MOVBU", "%s, %s+%d(FP)", reg, arg.Name, offset)
			case 2:
				buf.I("MOVHU", "%s, %s+%d(FP)", reg, arg.Name, offset)
			case 4:
				buf.I("MOVWU", "%s, %s+%d(FP)", reg, arg.Name, offset)
			case 8:
				buf.I("MOVD", "%s, %s+%d(FP)", reg, arg.Name, offset)
			default:
				panic(fmt.Sprintf("unknown int size: %d", size))
			}
		} else {
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
		}
	default:
		panic(fmt.Sprintf("unsupported return type: %T", arg.Type))
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
		offset = arch.loadArg(buf, &f.Args[i], offset)
	}

	seed := [32]byte{}
	copy(seed[:], f.Name)
	rnd := rand.New(rand.NewChaCha8(seed))

	// Set frame guard
	guardValue := rnd.Uint32()
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
