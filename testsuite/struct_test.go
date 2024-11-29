package testsuite

import (
	"math/rand/v2"
	"testing"

	"github.com/maxpoletaev/directcgo/testsuite/binding"
)

func compareResults(t *testing.T, result, expected []Pair) {
	t.Helper()

	if len(result) != len(expected) {
		t.Fatalf("expected %d pairs, got %d", len(expected), len(result))
	}

	for i, pair := range expected {
		if result[i] != pair {
			t.Fatalf("results do not match:\nwant: %v\ngot:  %v", expected, result)
		}
	}
}

func TestPassSmallStructSameIntegers(t *testing.T) {
	s := binding.SmallStructSameIntegers{
		U32_0: rand.Uint32(),
		U32_1: rand.Uint32(),
	}

	PassSmallStructSameIntegers(s)
	result := getOutput()

	PassSmallStructSameIntegersCgo(s)
	expected := getOutput()

	compareResults(t, result, expected)
}

func TestPassSmallStructMixedIntegers(t *testing.T) {
	s := binding.SmallStructMixedIntegers{
		U8:  uint8(rand.Uint32()),
		I32: int32(rand.Uint32()),
		U16: uint16(rand.Uint32()),
	}

	PassSmallStructMixedIntegers(s)
	result := getOutput()

	PassSmallStructMixedIntegersCgo(s)
	expected := getOutput()

	compareResults(t, result, expected)
}

func TestPassSmallStructSameFloats(t *testing.T) {
	s := binding.SmallStructSameFloats{
		F32_0: rand.Float32(),
		F32_1: rand.Float32(),
		F32_2: rand.Float32(),
	}

	PassSmallStructSameFloats(s)
	result := getOutput()

	PassSmallStructSameFloatsCgo(s)
	expected := getOutput()

	compareResults(t, result, expected)
}

func TestPassSmallStructMixedFloats(t *testing.T) {
	s := binding.SmallStructMixedFloats{
		F32_0: rand.Float32(),
		F64_0: rand.Float64(),
	}

	PassSmallStructMixedFloats(s)
	result := getOutput()

	PassSmallStructMixedFloatsCgo(s)
	expected := getOutput()

	compareResults(t, result, expected)
}

func TestPassSmallStructMixedNumbers(t *testing.T) {
	s := binding.SmallStructMixedNumbers{
		I32: int32(rand.Uint32()),
		U8:  uint8(rand.Uint32()),
		F32: rand.Float32(),
		U16: uint16(rand.Uint32()),
	}

	PassSmallStructMixedNumbers(s)
	result := getOutput()

	PassSmallStructMixedNumbersCgo(s)
	expected := getOutput()

	compareResults(t, result, expected)
}

func TestPassSmallStructNested(t *testing.T) {
	s := binding.SmallStructOuter{
		Inner_0: binding.SmallStructInner{U32: rand.Uint32()},
		Inner_1: binding.SmallStructInner{U32: rand.Uint32()},
	}

	PassSmallStructNested(s)
	result := getOutput()

	PassSmallStructNestedCgo(s)
	expected := getOutput()

	compareResults(t, result, expected)
}

func TestPassSmallStructWithArray(t *testing.T) {
	s := binding.SmallStructWithArray{
		Arr: [3]uint8{
			uint8(rand.Uint()),
			uint8(rand.Uint()),
			uint8(rand.Uint()),
		},
		U8:  uint8(rand.Uint()),
		F64: rand.Float64(),
	}

	PassSmallStructWithArray(s)
	result := getOutput()

	PassSmallStructWithArrayCgo(s)
	expected := getOutput()

	compareResults(t, result, expected)
}
