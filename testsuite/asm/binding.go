package asm

import "unsafe"

// void PassIntegers(int32_t i32, int64_t i64, int16_t i16, int8_t i8);
// void PassUnsignedIntegers(uint32_t u32, uint64_t u64, uint8_t u8, uint16_t u16);
// void PassFloats(float f32_0, double f64_0, double f64_1, float f32_1);
// void PassMixedNumbers(int8_t i8, float f32, uint32_t u32, double f64, int64_t i64);

//go:noescape
func PassIntegers(fn unsafe.Pointer, i32 int32, i64 int64, i16 int16, i8 int8)

//go:noescape
func PassUnsignedIntegers(fn unsafe.Pointer, u32 uint32, u64 uint64, u8 uint8, u16 uint16)

//go:noescape
func PassFloats(fn unsafe.Pointer, f32_0 float32, f64_0 float64, f64_1 float64, f32_1 float32)

//go:noescape
func PassMixedNumbers(fn unsafe.Pointer, i8 int8, f32 float32, u32 uint32, f64 float64, i64 int64)

//uint8_t ReturnUInt8(void);
//int8_t ReturnInt8(void);
//uint32_t ReturnUInt32(void);
//int32_t ReturnInt32(void);
//uint64_t ReturnUInt64(void);
//int64_t ReturnInt64(void);
//float ReturnFloat(void);
//double ReturnDouble(void);

//go:noescape
func ReturnUInt8(fn unsafe.Pointer) uint8

//go:noescape
func ReturnInt8(fn unsafe.Pointer) int8

//go:noescape
func ReturnUInt32(fn unsafe.Pointer) uint32

//go:noescape
func ReturnInt32(fn unsafe.Pointer) int32

//go:noescape
func ReturnUInt64(fn unsafe.Pointer) uint64

//go:noescape
func ReturnInt64(fn unsafe.Pointer) int64

//go:noescape
func ReturnFloat(fn unsafe.Pointer) float32

//go:noescape
func ReturnDouble(fn unsafe.Pointer) float64

//	typedef struct {
//	   uint8_t u8;
//	   int32_t i32;
//	   uint16_t u16;
//	} SmallStructIntegers;
type SmallStructIntegers struct {
	U8  uint8
	I32 int32
	U16 uint16
}

//	typedef struct {
//	   float f32;
//	   double f64;
//	} SmallStructFloats;
type SmallStructFloats struct {
	F32 float32
	F64 float64
}

//	typedef struct {
//	    int32_t i32;
//	    uint8_t u8;
//	    float f32;
//	    uint16_t u16;
//	} SmallStructMixed;
type SmallStructMixed struct {
	I32 int32
	U8  uint8
	F32 float32
	U16 uint16
}

// void PassSmallStructIntegers(SmallStructIntegers s);
// void PassSmallStructFloats(SmallStructFloats s);
// void PassSmallStructMixed(SmallStructMixed s);

//go:noescape
func PassSmallStructIntegers(fn unsafe.Pointer, s SmallStructIntegers)

//go:noescape
func PassSmallStructFloats(fn unsafe.Pointer, s SmallStructFloats)

//go:noescape
func PassSmallStructMixed(fn unsafe.Pointer, s SmallStructMixed)

//uint32_t AddTwoNumbers(uint32_t a, uint32_t b);

//go:noescape
func AddTwoNumbers(fn unsafe.Pointer, a uint32, b uint32) uint32
