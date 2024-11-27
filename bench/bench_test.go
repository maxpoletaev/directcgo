package bench

import (
	"math/rand/v2"
	"runtime"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	for i := 0; i < 1000; i++ {
		var (
			a = rand.Uint32()
			b = rand.Uint32()
		)

		want := AddTwoNumbersNative(a, b)
		got := AddTwoNumbersDirect(a, b)

		if got != want {
			t.Fatalf("want %d, got %d", want, got)
		}
	}
}

func BenchmarkAddTwoNumbersCgo(b *testing.B) {
	var res uint32
	x := rand.Uint32()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = AddTwoNumbersCgo(res, x)
	}

	runtime.KeepAlive(res)
}

func BenchmarkAddTwoNumbersDirect(b *testing.B) {
	var res uint32
	x := rand.Uint32()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = AddTwoNumbersDirect(res, x)
	}

	runtime.KeepAlive(res)
}

func BenchmarkAddTwoNumbersCodegen(b *testing.B) {
	var res uint32
	x := rand.Uint32()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = AddTwoNumbersCodegen(res, x)
	}

	runtime.KeepAlive(res)
}

func BenchmarkAddTwoNumbersNative(b *testing.B) {
	var res uint32
	x := rand.Uint32()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res = AddTwoNumbersNative(res, x)
	}

	runtime.KeepAlive(res)
}

func BenchmarkAddTwoNumbersLoopCgo(b *testing.B) {
	var res uint32
	x := rand.Uint32()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {
			res = AddTwoNumbersCgo(res, x)
		}
	}

	runtime.KeepAlive(res)
}

func BenchmarkAddTwoNumbersLoopDirect(b *testing.B) {
	var res uint32
	x := rand.Uint32()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {
			res = AddTwoNumbersDirect(res, x)
		}
	}

	runtime.KeepAlive(res)
}

func BenchmarkAddTwoNumbersLoopCodegen(b *testing.B) {
	var res uint32
	x := rand.Uint32()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {
			res = AddTwoNumbersCodegen(res, x)
		}
	}

	runtime.KeepAlive(res)
}

func BenchmarkAddTwoNumbersLoopNative(b *testing.B) {
	var res uint32
	x := rand.Uint32()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {
			res = AddTwoNumbersNative(res, x)
		}
	}

	runtime.KeepAlive(res)
}
