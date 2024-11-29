//go:build arm64

#include "textflag.h"

//func AddTwoNumbers(a, b uint32) uint32
TEXT ·AddTwoNumbers(SB), NOSPLIT, $0-12
    MOVWU    a+0(FP), R0
    MOVWU    b+4(FP), R1
    ADDW     R0, R1
    MOVWU    R1, ret+8(FP)
    RET
