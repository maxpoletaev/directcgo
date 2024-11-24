//go:build amd64

#include "go_asm.h"
#include "textflag.h"
#include "funcdata.h"
#include "directcgo.h"

// func call(fn unsafe.Pointer, arg unsafe.Pointer, ret unsafe.Pointer)
TEXT ·Call(SB), $FRAME_SIZE-24 // 1MB stack frame, 24 bytes for parameters
    NO_LOCAL_POINTERS
    MOVQ    fn+0(FP), AX
    MOVQ    arg+8(FP), DI
    MOVQ    ret+16(FP), SI
#ifdef FRAME_GUARD
    MOVL $0xDEADBEEF, 8(SP)
#endif
    MOVQ    SP, R12           // preserve original SP (callee-saved)
    LEAQ    FRAME_SIZE(SP), SP
    ANDQ    $~15, SP          // align to 16 bytes (ABI requirement)
    CALL    AX                // call C function
    MOVQ    R12, SP           // restore original SP
#ifdef FRAME_GUARD
    MOVL $0xDEADBEEF, AX
    CMPL AX, 8(SP)
    JNE overflow
#endif
    RET

overflow:
    CALL runtime·abort(SB)
    RET
