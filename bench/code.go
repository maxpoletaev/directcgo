package bench

/*
#include "code.h"
#include <stdint.h>

void ffi_AddTwoNumbers(void *arg, void *ret) {
	struct { uint32_t a; uint32_t b; } *a = arg;
	uint32_t sum = AddTwoNumbers(a->a, a->b);
	*(uint32_t *)ret = sum;
}
*/
import "C"
import (
	"unsafe"

	"github.com/maxpoletaev/directcgo"
	"github.com/maxpoletaev/directcgo/bench/asm"
)

func AddTwoNumbersDirect(a, b uint32) (ret uint32) {
	type args struct {
		a uint32
		b uint32
	}
	directcgo.Call(
		C.ffi_AddTwoNumbers,
		unsafe.Pointer(&args{a, b}),
		unsafe.Pointer(&ret),
	)
	return ret
}

func AddTwoNumbersCodegen(a, b uint32) uint32 {
	return asm.AddTwoNumbers(C.AddTwoNumbers, a, b)
}

func AddTwoNumbersCgo(a, b uint32) uint32 {
	return uint32(C.AddTwoNumbers(C.uint32_t(a), C.uint32_t(b)))
}

//go:noinline // we're measuring function call, not the addition instruction
func AddTwoNumbersNative(a, b uint32) uint32 {
	return a + b
}
