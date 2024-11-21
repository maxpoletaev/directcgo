package test

/*
#include "code.h"
#include <stdint.h>

void ffi_addTwoNumbers(void *arg, void *ret) {
	struct {
		uint32_t a;
		uint32_t b;
	} *a = arg;

	uint32_t result = addTwoNumbers(a->a, a->b);
	*(uint32_t *)ret = result;
}
*/
import "C"
import (
	"unsafe"

	"github.com/maxpoletaev/directcgo"
)

func AddTwoNumbersCgo(a, b uint32) uint32 {
	return uint32(C.addTwoNumbers(C.uint32_t(a), C.uint32_t(b)))
}

func AddTwoNumbersDirect(a, b uint32) (ret uint32) {
	type args struct {
		a uint32
		b uint32
	}

	directcgo.Call(C.ffi_addTwoNumbers, unsafe.Pointer(&args{a, b}), unsafe.Pointer(&ret))
	return ret
}

//go:noinline
func AddTwoNumbersNative(a, b uint32) uint32 {
	return a + b
}
