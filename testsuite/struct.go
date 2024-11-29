package testsuite

/*
#include "testsuite.h"
*/
import "C"
import "github.com/maxpoletaev/directcgo/testsuite/binding"

/*
void PassSmallStructSameIntegers(SmallStructSameIntegers s);
void PassSmallStructMixedIntegers(SmallStructMixedIntegers s);
void PassSmallStructSameFloats(SmallStructSameFloats s);
void PassSmallStructMixedFloats(SmallStructMixedFloats s);
void PassSmallStructMixedNumbers(SmallStructMixedNumbers s);
void PassSmallStructNested(SmallStructOuter s);
*/

func PassSmallStructSameIntegers(s binding.SmallStructSameIntegers) {
	binding.PassSmallStructSameIntegers(C.PassSmallStructSameIntegers, s)
}

func PassSmallStructSameIntegersCgo(s binding.SmallStructSameIntegers) {
	C.PassSmallStructSameIntegers(C.SmallStructSameIntegers{
		u32_0: C.uint32_t(s.U32_0),
		u32_1: C.uint32_t(s.U32_1),
	})
}

func PassSmallStructMixedIntegers(s binding.SmallStructMixedIntegers) {
	binding.PassSmallStructMixedIntegers(C.PassSmallStructMixedIntegers, s)
}

func PassSmallStructMixedIntegersCgo(s binding.SmallStructMixedIntegers) {
	C.PassSmallStructMixedIntegers(C.SmallStructMixedIntegers{
		u8:  C.uint8_t(s.U8),
		i32: C.int32_t(s.I32),
		u16: C.uint16_t(s.U16),
	})
}

func PassSmallStructSameFloats(s binding.SmallStructSameFloats) {
	binding.PassSmallStructSameFloats(C.PassSmallStructSameFloats, s)
}

func PassSmallStructSameFloatsCgo(s binding.SmallStructSameFloats) {
	C.PassSmallStructSameFloats(C.SmallStructSameFloats{
		f32_0: C.float(s.F32_0),
		f32_1: C.float(s.F32_1),
		f32_2: C.float(s.F32_2),
	})
}

func PassSmallStructMixedFloats(s binding.SmallStructMixedFloats) {
	binding.PassSmallStructMixedFloats(C.PassSmallStructMixedFloats, s)
}

func PassSmallStructMixedFloatsCgo(s binding.SmallStructMixedFloats) {
	C.PassSmallStructMixedFloats(C.SmallStructMixedFloats{
		f32_0: C.float(s.F32_0),
		f64_0: C.double(s.F64_0),
	})
}

func PassSmallStructMixedNumbers(s binding.SmallStructMixedNumbers) {
	binding.PassSmallStructMixedNumbers(C.PassSmallStructMixedNumbers, s)
}

func PassSmallStructMixedNumbersCgo(s binding.SmallStructMixedNumbers) {
	C.PassSmallStructMixedNumbers(C.SmallStructMixedNumbers{
		i32: C.int32_t(s.I32),
		u8:  C.uint8_t(s.U8),
		f32: C.float(s.F32),
		u16: C.uint16_t(s.U16),
	})
}

func PassSmallStructNested(s binding.SmallStructOuter) {
	binding.PassSmallStructNested(C.PassSmallStructNested, s)
}

func PassSmallStructNestedCgo(s binding.SmallStructOuter) {
	C.PassSmallStructNested(C.SmallStructOuter{
		inner_0: C.SmallStructInner{
			u32: C.uint32_t(s.Inner_0.U32),
		},
		inner_1: C.SmallStructInner{
			u32: C.uint32_t(s.Inner_1.U32),
		},
	})
}
