package testsuite

import (
	"runtime"
	"testing"

	"math/rand/v2"
)

//go:noinline
func addTwoNumbersNative(a, b uint32) uint32 {
	return a + b
}

func BenchmarkAddTwoNumbers(b *testing.B) {
	x := rand.Uint32()

	b.Run("CGO", func(b *testing.B) {
		var ret uint32
		for i := 0; i < b.N; i++ {
			ret = AddTwoNumbersCgo(ret, x)
		}
		runtime.KeepAlive(ret)
	})

	b.Run("Direct", func(b *testing.B) {
		var ret uint32
		for i := 0; i < b.N; i++ {
			ret = AddTwoNumbers(ret, x)
		}
		runtime.KeepAlive(ret)
	})

	b.Run("Native", func(b *testing.B) {
		var ret uint32
		for i := 0; i < b.N; i++ {
			ret = addTwoNumbersNative(ret, x)
		}
		runtime.KeepAlive(ret)
	})
}
