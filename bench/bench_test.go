package bench

import (
	"math/rand/v2"
	"runtime"
	"testing"

	"github.com/maxpoletaev/directcgo/bench/pureasm"
)

func TestAddTwoNumbers(t *testing.T) {
	for i := 0; i < 1000; i++ {
		var (
			a = rand.Uint32()
			b = rand.Uint32()
		)

		directCallResult := AddTwoNumbersDirectCall(a, b)
		pureasmResult := pureasm.AddTwoNumbers(a, b)
		want := AddTwoNumbersNative(a, b)

		if directCallResult != want || pureasmResult != want {
			t.Fatalf("want %d, got [%d, %d]", want, directCallResult, pureasmResult)
		}
	}
}

func BenchmarkAddTwoNumbers(b *testing.B) {
	b.Run("cgo", func(b *testing.B) {
		var res uint32
		x := rand.Uint32()
		for i := 0; i < b.N; i++ {
			res = AddTwoNumbersCgo(res, x)
		}
		runtime.KeepAlive(res)
	})

	b.Run("directcgo", func(b *testing.B) {
		var res uint32
		x := rand.Uint32()
		for i := 0; i < b.N; i++ {
			res = AddTwoNumbersDirectCall(res, x)
		}
		runtime.KeepAlive(res)
	})

	b.Run("codegen", func(b *testing.B) {
		var res uint32
		x := rand.Uint32()
		for i := 0; i < b.N; i++ {
			res = AddTwoNumbersCodegen(res, x)
		}
		runtime.KeepAlive(res)
	})

	b.Run("pureasm", func(b *testing.B) {
		var res uint32
		x := rand.Uint32()
		for i := 0; i < b.N; i++ {
			res = pureasm.AddTwoNumbers(res, x)
		}
		runtime.KeepAlive(res)
	})

	b.Run("native", func(b *testing.B) {
		var res uint32
		x := rand.Uint32()
		for i := 0; i < b.N; i++ {
			res = AddTwoNumbersNative(res, x)
		}
		runtime.KeepAlive(res)
	})
}

func BenchmarkAddTwoNumbersLoop100(b *testing.B) {
	const iterations = 100

	b.Run("cgo", func(b *testing.B) {
		var res uint32
		x := rand.Uint32()
		for i := 0; i < b.N; i++ {
			for j := 0; j < iterations; j++ {
				res = AddTwoNumbersCgo(res, x)
			}
		}
		runtime.KeepAlive(res)
	})

	b.Run("directcgo", func(b *testing.B) {
		var res uint32
		x := rand.Uint32()
		for i := 0; i < b.N; i++ {
			for j := 0; j < iterations; j++ {
				res = AddTwoNumbersDirectCall(res, x)
			}
		}
		runtime.KeepAlive(res)
	})

	b.Run("codegen", func(b *testing.B) {
		var res uint32
		x := rand.Uint32()
		for i := 0; i < b.N; i++ {
			for j := 0; j < iterations; j++ {
				res = AddTwoNumbersCodegen(res, x)
			}
		}
		runtime.KeepAlive(res)
	})

	b.Run("pureasm", func(b *testing.B) {
		var res uint32
		x := rand.Uint32()
		for i := 0; i < b.N; i++ {
			for j := 0; j < iterations; j++ {
				res = pureasm.AddTwoNumbers(res, x)
			}
		}
		runtime.KeepAlive(res)
	})

	b.Run("native", func(b *testing.B) {
		var res uint32
		x := rand.Uint32()
		for i := 0; i < b.N; i++ {
			for j := 0; j < iterations; j++ {
				res = AddTwoNumbersNative(res, x)
			}
		}
		runtime.KeepAlive(res)
	})
}
