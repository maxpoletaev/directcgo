package main

/*
#include "code.h"
#include <stdint.h>
#include <stdlib.h>

void ffi_PrintHello(void *arg, void *ret)
{
	struct {
		char *name;
	} *a = arg;

	PrintHello(a->name);
}

void ffi_AddTwoNumbers(void *arg, void *ret)
{
	struct {
		uint32_t a;
		uint32_t b;
	} *a = arg;

	uint32_t sum = AddTwoNumbers(a->a, a->b);
	*(uint32_t *)ret = sum;
}
*/
import "C"
import (
	"fmt"
	"math/rand"
	"unsafe"

	"github.com/maxpoletaev/directcgo"
)

func addTwoNumbers(a, b uint32) (ret uint32) {
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

func printHello(name string) {
	type args struct {
		name *C.char
	}

	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	directcgo.Call(
		C.ffi_PrintHello,
		unsafe.Pointer(&args{cName}),
		nil,
	)
}

func main() {
	var (
		a = uint32(rand.Intn(100))
		b = uint32(rand.Intn(100))
	)
	printHello("directcgo")
	sum := addTwoNumbers(a, b)
	fmt.Printf("%d + %d = %d\n", a, b, sum)
}
