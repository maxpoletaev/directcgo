//go:build amd64 && !windows

// Code generated by directcgo. DO NOT EDIT.
// directcgo -arch=amd64,arm64 ./bench/asm

#include "go_asm.h"
#include "textflag.h"
#include "funcdata.h"

// AddTwoNumbers func(fn unsafe.Pointer, a, b uint32) uint32
TEXT ·AddTwoNumbers(SB), $65536-20
	MOVQ    fn+0(FP), AX
	MOVLQZX a+8(FP), DI
	MOVLQZX b+12(FP), SI
	MOVL    $0x841709CB, R10
	MOVL    R10, 8(SP)
	MOVQ    SP, R12
	LEAQ    65536(SP), SP
	ANDQ    $~15, SP
	CALL    AX
	MOVQ    R12, SP
	MOVL    8(SP), R10
	CMPL    R10, $0x841709CB
	JNE     overflow
	MOVL    AX, ret+16(FP)
	RET
overflow:
	CALL    runtime·abort(SB)
	RET