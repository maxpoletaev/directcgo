package binding

import "unsafe"

// -----------------------
// Passing primitive types
// -----------------------

/*
void PassIntegers(int32_t i32, int64_t i64, int16_t i16, int8_t i8);
void PassUnsignedIntegers(uint32_t u32, uint64_t u64, uint8_t u8, uint16_t u16);
void PassFloats(float f32_0, double f64_0, double f64_1, float f32_1);
void PassMixedNumbers(int8_t i8, float f32, uint32_t u32, double f64, int64_t i64);
*/

//go:noescape
func PassIntegers(fn unsafe.Pointer, i32 int32, i64 int64, i16 int16, i8 int8)

//go:noescape
func PassUnsignedIntegers(fn unsafe.Pointer, u32 uint32, u64 uint64, u8 uint8, u16 uint16)

//go:noescape
func PassFloats(fn unsafe.Pointer, f32_0 float32, f64_0 float64, f64_1 float64, f32_1 float32)

//go:noescape
func PassMixedNumbers(fn unsafe.Pointer, i8 int8, f32 float32, u32 uint32, f64 float64, i64 int64)

// -------------------------
// Returning primitive types
// -------------------------

/*
uint8_t ReturnUInt8(void);
int8_t ReturnInt8(void);
uint32_t ReturnUInt32(void);
int32_t ReturnInt32(void);
uint64_t ReturnUInt64(void);
int64_t ReturnInt64(void);
float ReturnFloat(void);
double ReturnDouble(void);
*/

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

// ---------------
// Passing structs
// ---------------

//	typedef struct {
//	   uint32_t u32_0;
//	   uint32_t u32_1;
//	} SmallStructSameIntegers;
type SmallStructSameIntegers struct {
	U32_0 uint32
	U32_1 uint32
}

//	typedef struct {
//	   uint8_t u8;
//	   int32_t i32;
//	   uint16_t u16;
//	} SmallStructMixedIntegers;
type SmallStructMixedIntegers struct {
	U8  uint8
	I32 int32
	U16 uint16
}

//	typedef struct {
//	   float f32_0;
//	   float f32_1;
//	   float f32_2;
//	} SmallStructSameFloats;
type SmallStructSameFloats struct {
	F32_0 float32
	F32_1 float32
	F32_2 float32
}

//	typedef struct {
//	   float f32_0;
//	   double f64_0;
//	} SmallStructMixedFloats;
type SmallStructMixedFloats struct {
	F32_0 float32
	F64_0 float64
}

//	typedef struct {
//	   int32_t i32;
//	   uint8_t u8;
//	   float f32;
//	   uint16_t u16;
//	} SmallStructMixedNumbers;
type SmallStructMixedNumbers struct {
	I32 int32
	U8  uint8
	F32 float32
	U16 uint16
}

//	typedef struct {
//	   uint32_t u32;
//	} SmallStructInner;
type SmallStructInner struct {
	U32 uint32
}

//	typedef struct {
//	   SmallStructInner inner_0;
//	   SmallStructInner inner_1;
//	} SmallStructOuter;
type SmallStructOuter struct {
	Inner_0 SmallStructInner
	Inner_1 SmallStructInner
}

//	typedef struct {
//	   uint8_t u8;
//	   uint8_t arr[3];
//	   double f64;
//	} SmallStructWithArray;
type SmallStructWithArray struct {
	U8  uint8
	Arr [3]uint8
	F64 float64
}

//	typedef struct {
//	   uint32_t u32_0;
//	   uint32_t u32_1;
//	   uint32_t u32_2;
//	   uint32_t u32_3;
//	   uint32_t u32_4;
//	   uint32_t u32_5;
//	   uint32_t u32_6;
//	   uint32_t u32_7;
//	   uint32_t u32_8;
//	   uint32_t u32_9;
//	} LargeStruct;
type LargeStruct struct {
	U32_0 uint32
	U32_1 uint32
	U32_2 uint32
	U32_3 uint32
	U32_4 uint32
	U32_5 uint32
	U32_6 uint32
	U32_7 uint32
	U32_8 uint32
	U32_9 uint32
}

/*
void PassSmallStructSameIntegers(SmallStructSameIntegers s);
void PassSmallStructMixedIntegers(SmallStructMixedIntegers s);
void PassSmallStructSameFloats(SmallStructSameFloats s);
void PassSmallStructMixedFloats(SmallStructMixedFloats s);
void PassSmallStructMixedNumbers(SmallStructMixedNumbers s);
void PassSmallStructNested(SmallStructOuter s);
void PassSmallStructWithArray(SmallStructWithArray s);
void PassLargeStruct(LargeStruct s);
*/

//go:noescape
func PassSmallStructSameIntegers(fn unsafe.Pointer, s SmallStructSameIntegers)

//go:noescape
func PassSmallStructMixedIntegers(fn unsafe.Pointer, s SmallStructMixedIntegers)

//go:noescape
func PassSmallStructSameFloats(fn unsafe.Pointer, s SmallStructSameFloats)

//go:noescape
func PassSmallStructMixedFloats(fn unsafe.Pointer, s SmallStructMixedFloats)

//go:noescape
func PassSmallStructMixedNumbers(fn unsafe.Pointer, s SmallStructMixedNumbers)

//go:noescape
func PassSmallStructNested(fn unsafe.Pointer, s SmallStructOuter)

//go:noescape
func PassSmallStructWithArray(fn unsafe.Pointer, s SmallStructWithArray)

//go:noescape
func PassLargeStruct(fn unsafe.Pointer, s LargeStruct)
