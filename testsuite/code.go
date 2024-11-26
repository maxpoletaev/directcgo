package testsuite

/*
#include "code.h"
*/
import "C"
import "github.com/maxpoletaev/directcgo/testsuite/asm"

// void PassIntegers(int32_t i32, int64_t i64, int16_t i16, int8_t i8);
// void PassUnsignedIntegers(uint32_t u32, uint64_t u64, uint8_t u8, uint16_t u16);
// void PassFloats(float f32_0, double f64_0, double f64_1, float f32_1);
// void PassMixedNumbers(int8_t i8, float f32, uint32_t u32, double f64, int64_t i64);

func PassIntegers(i32 int32, i64 int64, i16 int16, i8 int8) {
	asm.PassIntegers(C.PassIntegers, i32, i64, i16, i8)
}

func PassIntegersCgo(i32 int32, i64 int64, i16 int16, i8 int8) {
	C.PassIntegers(C.int32_t(i32), C.int64_t(i64), C.int16_t(i16), C.int8_t(i8))
}

func PassUnsignedIntegers(u32 uint32, u64 uint64, u8 uint8, u16 uint16) {
	asm.PassUnsignedIntegers(C.PassUnsignedIntegers, u32, u64, u8, u16)
}

func PassUnsignedIntegersCgo(u32 uint32, u64 uint64, u8 uint8, u16 uint16) {
	C.PassUnsignedIntegers(C.uint32_t(u32), C.uint64_t(u64), C.uint8_t(u8), C.uint16_t(u16))
}

func PassFloats(f32_0 float32, f64_0 float64, f64_1 float64, f32_1 float32) {
	asm.PassFloats(C.PassFloats, f32_0, f64_0, f64_1, f32_1)
}

func PassFloatsCgo(f32_0 float32, f64_0 float64, f64_1 float64, f32_1 float32) {
	C.PassFloats(C.float(f32_0), C.double(f64_0), C.double(f64_1), C.float(f32_1))
}

func PassMixedNumbers(i8 int8, f32 float32, u32 uint32, f64 float64, i64 int64) {
	asm.PassMixedNumbers(C.PassMixedNumbers, i8, f32, u32, f64, i64)
}

func PassMixedNumbersCgo(i8 int8, f32 float32, u32 uint32, f64 float64, i64 int64) {
	C.PassMixedNumbers(C.int8_t(i8), C.float(f32), C.uint32_t(u32), C.double(f64), C.int64_t(i64))
}

//uint8_t ReturnUInt8(void);
//int8_t ReturnInt8(void);
//uint32_t ReturnUInt32(void);
//int32_t ReturnInt32(void);
//uint64_t ReturnUInt64(void);
//int64_t ReturnInt64(void);
//float ReturnFloat(void);
//double ReturnDouble(void);

func ReturnUInt8() uint8 {
	return asm.ReturnUInt8(C.ReturnUInt8)
}

func ReturnUInt8Cgo() uint8 {
	return uint8(C.ReturnUInt8())
}

func ReturnInt8() int8 {
	return asm.ReturnInt8(C.ReturnInt8)
}

func ReturnInt8Cgo() int8 {
	return int8(C.ReturnInt8())
}

func ReturnUInt32() uint32 {
	return asm.ReturnUInt32(C.ReturnUInt32)
}

func ReturnUInt32Cgo() uint32 {
	return uint32(C.ReturnUInt32())
}

func ReturnInt32() int32 {
	return asm.ReturnInt32(C.ReturnInt32)
}

func ReturnInt32Cgo() int32 {
	return int32(C.ReturnInt32())
}

func ReturnUInt64() uint64 {
	return asm.ReturnUInt64(C.ReturnUInt64)
}

func ReturnUInt64Cgo() uint64 {
	return uint64(C.ReturnUInt64())
}

func ReturnInt64() int64 {
	return asm.ReturnInt64(C.ReturnInt64)
}

func ReturnInt64Cgo() int64 {
	return int64(C.ReturnInt64())
}

func ReturnFloat() float32 {
	return asm.ReturnFloat(C.ReturnFloat)
}

func ReturnFloatCgo() float32 {
	return float32(C.ReturnFloat())
}

func ReturnDouble() float64 {
	return asm.ReturnDouble(C.ReturnDouble)
}

func ReturnDoubleCgo() float64 {
	return float64(C.ReturnDouble())
}

// void PassSmallStructIntegers(SmallStructIntegers s);
// void PassSmallStructFloats(SmallStructFloats s);
// void PassSmallStructMixed(SmallStructMixed s);

func PassSmallStructIntegers(s asm.SmallStructIntegers) {
	asm.PassSmallStructIntegers(C.PassSmallStructIntegers, s)
}

func PassSmallStructIntegersCgo(s asm.SmallStructIntegers) {
	C.PassSmallStructIntegers(C.SmallStructIntegers{
		u8:  C.uint8_t(s.U8),
		i32: C.int32_t(s.I32),
	})
}

func PassSmallStructFloats(s asm.SmallStructFloats) {
	asm.PassSmallStructFloats(C.PassSmallStructFloats, s)
}

func PassSmallStructFloatsCgo(s asm.SmallStructFloats) {
	C.PassSmallStructFloats(C.SmallStructFloats{
		f32: C.float(s.F32),
		f64: C.double(s.F64),
	})
}

func PassSmallStructMixed(s asm.SmallStructMixed) {
	asm.PassSmallStructMixed(C.PassSmallStructMixed, s)
}

func PassSmallStructMixedCgo(s asm.SmallStructMixed) {
	C.PassSmallStructMixed(C.SmallStructMixed{
		i32: C.int32_t(s.I32),
		u8:  C.uint8_t(s.U8),
		f32: C.float(s.F32),
		u16: C.uint16_t(s.U16),
	})
}

// uint32_t AddTwoNumbers(uint32_t a, uint32_t b);

func AddTwoNumbers(a uint32, b uint32) uint32 {
	return asm.AddTwoNumbers(C.AddTwoNumbers, a, b)
}

func AddTwoNumbersCgo(a uint32, b uint32) uint32 {
	return uint32(C.AddTwoNumbers(C.uint32_t(a), C.uint32_t(b)))
}
