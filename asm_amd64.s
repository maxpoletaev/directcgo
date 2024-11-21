//go:build amd64

#include "go_asm.h"
#include "textflag.h"

// func call(fn unsafe.Pointer, arg unsafe.Pointer, ret unsafe.Pointer)
TEXT ·Call(SB), $1048576-24 // 1MB stack frame, 24 bytes for parameters
    MOVQ    fn+0(FP), AX
    MOVQ    arg+8(FP), DI
    MOVQ    ret+16(FP), SI
    MOVQ    SP, DX          // Save original stack pointer
    ANDQ    $~15, SP        // Align stack to 16 bytes (clear bottom 4 bits)
    MOVQ    DX, 8(SP)       // Save old SP for restoration
    CALL    AX              // Call the function
    MOVQ    8(SP), DX       // Restore saved SP
    MOVQ    DX, SP          // Restore the stack pointer
    RET
