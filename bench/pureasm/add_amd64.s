//go:build amd64

#include "textflag.h"

//func AddTwoNumbers(a, b uint32) uint32
TEXT ·AddTwoNumbers(SB), NOSPLIT, $0-12
    MOVL    a+0(FP), AX
    MOVL    b+4(FP), CX
    ADDL    CX, AX
    MOVL    AX, ret+8(FP)
    RET
