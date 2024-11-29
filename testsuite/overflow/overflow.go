package main

/*
#cgo CFLAGS: -O0

#include <stdio.h>
#include <string.h>

static void printStackPointer()
{
	void* sp;
#if defined(__x86_64__)
	__asm__("movq %%rsp, %0" : "=r"(sp));
#elif defined(__aarch64__)
	__asm__("mov %0, sp" : "=r"(sp));
#endif
	printf("SP: %p\n", sp);
}

void recursiveFunc(void *arg)
{
	struct {
		unsigned int n;
	} *a = arg;

	char buf[1024];
	memset(buf, 0, 1024);

	if (a->n == 0) {
		return;
	}

	a->n--;
	recursiveFunc(a);
}
*/
import "C"
import (
	"flag"
	"fmt"
	"log"
	"math/rand/v2"
	"runtime"
	"sync"
	"unsafe"

	"github.com/maxpoletaev/directcgo"
)

func main() {
	var (
		mode        int
		depth       int
		concurrency int
		iterations  int
		silent      bool
	)

	flag.IntVar(&mode, "mode", 0, "0: directcgo (default), 2: cgo")
	flag.IntVar(&concurrency, "concurrency", 1, "concurrency")
	flag.IntVar(&iterations, "iterations", 1, "number of iterations (0 = infinite)")
	flag.BoolVar(&silent, "silent", false, "disable printing")
	flag.IntVar(&depth, "depth", 10, "recursion depth")
	flag.Parse()

	type funcArgs struct {
		n uint32
	}

	wg := sync.WaitGroup{}
	wg.Add(concurrency)

	for i := 0; i < concurrency; i++ {
		go func() {
			defer wg.Done()

			for j := 0; j < iterations; j++ {
				for k := 0; k < depth; k++ {
					n := uint32(k)

					if !silent {
						fmt.Println("depth=", n)
					}

					// Trigger GC occasionally
					if rand.Float32() < 0.3 {
						runtime.GC()
					}

					switch mode {
					case 0:
						directcgo.Call(C.recursiveFunc, unsafe.Pointer(&funcArgs{n}), nil)
					case 1:
						C.recursiveFunc(unsafe.Pointer(&funcArgs{n}))
					default:
						log.Fatalf("invalid mode: %d", mode)
					}
				}
			}
		}()
	}

	wg.Wait()
}
