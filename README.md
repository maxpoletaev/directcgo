# directcgo - Minimal Cost C Interop Experiment

This is an experiment which targets one specific case: minimizing cost of calling fast short-lived C functions in a tight loop (e.g., immediate mode graphic apps, game engines, SIMD code). Basically, what the original cgo design is infamous for.

There are two main reasons why cgo is considered slow in such scenarios (of course, slow here [is not technically slow][4], we're still talking about dozens of nanoseconds of overhead per call, it's just that this overhead gets accumulated quickly when a function is called hundreds of times per frame). The technical aspects have probably been described and discussed in great detail across many posts and articles on the internet, but in a nutshell:

 1. Goroutines use dynamic stacks managed by the runtime. The stacks may grow, shrink and moved around at any time, which makes them not very suitable for C calls, nor is it possible to control the stack size from the C code being called. Due to this limitation, every cgo call is implemented by jumping to the system stack (the stack where the runtime is running), executing the code there, and jumping back.

 2. Go wants C calls to natively fit into Go's concurrency model, making an assumption that any cgo call might block indefinitely. Therefore, any C call requires additional cooperation with the scheduler, like excluding threads occupied with C code from the thread pool and spawning new threads to compensate for the loss. This does have certain benefits if the program mainly uses C functions to offload heavy work or I/O, but greatly reduces throughput for small and quick C functions which need to be called at higher frequency.

## How it works

The idea initially came from [fastcgo][1] (which, in turn, was inspired by [rustgo][2]) that avoids interacting with the scheduler and jumps to the system stack directly. A [thread][3] on Google Groups made me think if we could also get away with running C code right on the goroutine stack to avoid rapid context switches. And it turns out, we could!

What the code in this repo tries to do is basically run C code on the same goroutine stack, ensuring that there is a reasonable stack space available to safely run this code. This brings it very close to the efficiency of native function calls (because it is essentially calling native functions, using the standard ABI calling convention for the target architecture). Performance-wise, this has only 2.5-3x overhead (introduced by an extra layer of indirection through the assembly code) compared to native Go calls, while traditional Cgo slaps you with a 30-40x penalty.

```
BenchmarkAddTwoNumbersCgo-6          	36542678	        32.94 ns/op
BenchmarkAddTwoNumbersDirect-6       	419276713	         2.862 ns/op
BenchmarkAddTwoNumbersNative-6       	1000000000	         1.070 ns/op
BenchmarkAddTwoNumbersLoopCgo-6      	  362950	      3296 ns/op
BenchmarkAddTwoNumbersLoopDirect-6   	 4078764	       294.2 ns/op
BenchmarkAddTwoNumbersLoopNative-6   	10635810	       112.7 ns/op
```

The solution itself is so simple it's almost embarrassing, and pretty much future-proof, as it does not rely on any runtime internals. All we need is a so-called assembly trampoline that will set up a stack frame large enough to fit the entire C stack and arrange the arguments according to the ABI standard on the target platform:

```
//go:build arm64

// func Call(fn unsafe.Pointer, arg unsafe.Pointer, ret unsafe.Pointer)
TEXT ·Call(SB), $1048576-24 // 1MB stack frame, 24 bytes for parameters
    MOVD    fn+0(FP), R2
    MOVD    arg+8(FP), R0
    MOVD    ret+16(FP), R1
    MOVD    $1048576, R10
    ADD     R10, RSP
    BL      (R2)
    RET
```

And some wrapping code for passing function arguments and return values:

```go
/*
#include "code.h"
#include <stdint.h>

void ffi_addTwoNumbers(void *arg, void *ret) {
    struct { 
        uint32_t a; 
        uint32_t b;
    } *a = arg;

    uint32_t sum = addTwoNumbers(a->a, a->b);
    *(uint32_t *)ret = sum;
}
*/
import "C"
import "unsafe"

func AddTwoNumbers(a, b uint32) (ret uint32) {
    type args struct {
        a uint32
        b uint32
    }
    directcgo.Call(
        C.ffi_addTwoNumbers, 
        unsafe.Pointer(&args{a, b}), 
        unsafe.Pointer(&ret),
    )
    return ret
}
```

## If you want to use it

Unlike other packages, this one actually encourages you to copy its contents into your project instead of installing it as an external dependency, since you may need to tune the stack size (see caveats). Configuring it externally is not possible, since these are just compile-time preprocessor parameters for the assembly code generation.

There are also some basic stack frame overflow guards that can be enabled or disabled depending on how confident you are about the code you are calling (these are not very reliable, thus they are optional. The overflow can only be detected if the callee returns before it screws everything up).

Do extensive stress testing under various conditions and input sizes before seriously considering using directcgo instead of cgo. Ideally, you should know what the code is doing, how much stack space it actually needs and if there are any VLAs, alloca(), or recursion involved, that may allocate a lot of space on the stack and blow it up.

## Caveats

However, trying to bypass cgo in order to squeeze out every last bit of performance needs a good justification, as it comes with a lot of potential caveats:

 1. Unusual stack sizes. This method will create goroutines with stacks that are probably too large for spawning thousands of them, but not large enough to match the size of the system stack. This repo experiments with 1MB stack size, so prior to the C call, Go will allocate a 1MB stack frame for each goroutine calling to C code. That will likely lead to OOM when spawned uncontrollably, so you may want to create a dedicated pool of "fat" goroutines with bigger stacks for calling C code this way.

 2. C calls do not inform the scheduler about their presence. The thread calling to C code will block until the C function returns, effectively blocking other goroutines waiting in the run queue on that thread. There is no way for the scheduler to preempt the running C function, which, in turn, may lead to other negative side effects like delaying garbage collector phases due to the goroutine not being able to reach a safepoint. That's why I said the function should be short-lived.

 3. There aren't any security measures in place. Like, no cgocheck or anything like that for checking that you are passing pointers to pinned Go memory. Also, stack overflow of a goroutine stack is not guaranteed to cause the program crash, as it may overflow into valid memory reserved by the go allocator. So, no safety nets, no helmet, just good old undefined behavior.

It might be a good idea to use this method only for short-lived C functions in the hot path of the application, where the performance gain is significant enough to justify the risks. For everything else, just use cgo.

[1]: https://github.com/petermattis/fastcgo
[2]: https://words.filippo.io/rustgo/
[3]: https://groups.google.com/g/golang-nuts/c/_YrvM8OO6QY
[4]: https://shane.ai/posts/cgo-performance-in-go1.21/