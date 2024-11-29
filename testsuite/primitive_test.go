package testsuite

import (
	"math/rand/v2"
	"testing"
)

func TestPassIntegers(t *testing.T) {
	for i := 0; i < 100; i++ {
		i32 := rand.Int32()
		i64 := rand.Int64()
		i16 := int16(rand.Int())
		i8 := int8(rand.Int())

		PassIntegers(i32, i64, i16, i8)
		result := getOutput()
		PassIntegersCgo(i32, i64, i16, i8)
		expected := getOutput()

		compareResults(t, result, expected)
	}
}

func TestPassUnsignedIntegers(t *testing.T) {
	for i := 0; i < 100; i++ {
		u32 := rand.Uint32()
		u64 := rand.Uint64()
		u8 := uint8(rand.Int())
		u16 := uint16(rand.Int())

		PassUnsignedIntegers(u32, u64, u8, u16)
		result := getOutput()
		PassUnsignedIntegersCgo(u32, u64, u8, u16)
		expected := getOutput()

		compareResults(t, result, expected)
	}
}

func TestPassFloats(t *testing.T) {
	for i := 0; i < 100; i++ {
		f32_0 := rand.Float32()
		f64_0 := rand.Float64()
		f64_1 := rand.Float64()
		f32_1 := rand.Float32()

		PassFloats(f32_0, f64_0, f64_1, f32_1)
		result := getOutput()
		PassFloatsCgo(f32_0, f64_0, f64_1, f32_1)
		expected := getOutput()

		compareResults(t, result, expected)
	}
}

func TestPassMixedNumbers(t *testing.T) {
	for i := 0; i < 100; i++ {
		i8 := int8(rand.Int())
		f32 := rand.Float32()
		u32 := rand.Uint32()
		f64 := rand.Float64()
		i64 := rand.Int64()

		PassMixedNumbers(i8, f32, u32, f64, i64)
		result := getOutput()
		PassMixedNumbersCgo(i8, f32, u32, f64, i64)
		expected := getOutput()

		compareResults(t, result, expected)
	}
}

func TestReturnPrimitives(t *testing.T) {
	if ReturnUInt8() != ReturnUInt8Cgo() {
		t.Fatalf("ReturnUInt8 results do not match")
	}
	if ReturnInt8() != ReturnInt8Cgo() {
		t.Fatalf("ReturnInt8 results do not match")
	}
	if ReturnUInt32() != ReturnUInt32Cgo() {
		t.Fatalf("ReturnUInt32 results do not match")
	}
	if ReturnInt32() != ReturnInt32Cgo() {
		t.Fatalf("ReturnInt32 results do not match")
	}
	if ReturnUInt64() != ReturnUInt64Cgo() {
		t.Fatalf("ReturnUInt64 results do not match")
	}
	if ReturnInt64() != ReturnInt64Cgo() {
		t.Fatalf("ReturnInt64 results do not match")
	}
	if ReturnFloat() != ReturnFloatCgo() {
		t.Fatalf("ReturnFloat results do not match")
	}
	if ReturnDouble() != ReturnDoubleCgo() {
		t.Fatalf("ReturnDouble results do not match")
	}
}
