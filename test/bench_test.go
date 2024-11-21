package test

import (
	"math/rand"
	"runtime"
	"testing"
)

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
