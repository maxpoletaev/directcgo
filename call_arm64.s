//go:build arm64

#include "go_asm.h"
#include "textflag.h"
#include "funcdata.h"
#include "directcgo.h"

// func call(fn unsafe.Pointer, arg unsafe.Pointer, ret unsafe.Pointer)
TEXT ·Call(SB), $FRAME_SIZE-24 // reserve 1MB stack frame, 24 bytes for parameters
    NO_LOCAL_POINTERS
    MOVD    fn+0(FP), R9
    MOVD    arg+8(FP), R0
    MOVD    ret+16(FP), R1
    MOVD    RSP, R19           // preserve original SP (callee-saved)
    MOVD    RSP, R10
    AND     $~15, R10, RSP     // align to 16 bytes (ABI requirement)
    BL      (R9)               // call function
    MOVD    R19, RSP           // restore original SP
    RET
