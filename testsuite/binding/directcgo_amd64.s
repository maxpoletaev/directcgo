//go:build amd64 && !windows

// Code generated by directcgo. DO NOT EDIT.
// directcgo -arch=amd64,arm64 ./testsuite/binding

#include "go_asm.h"
#include "textflag.h"
#include "funcdata.h"

// PassIntegers func(fn unsafe.Pointer, i32 int32, i64 int64, i16 int16, i8 int8)
TEXT ·PassIntegers(SB), $65536-27
	MOVQ    fn+0(FP), AX
	MOVLQSX i32+8(FP), DI
	MOVQ    i64+16(FP), SI
	MOVWQSX i16+24(FP), DX
	MOVBQSX i8+26(FP), CX
	MOVL    $0xC14F9D00, R10
	MOVL    R10, 8(SP)
	MOVQ    SP, R12
	LEAQ    65536(SP), SP
	ANDQ    $~15, SP
	CALL    AX
	MOVQ    R12, SP
	MOVL    8(SP), R10
	CMPL    R10, $0xC14F9D00
	JNE     overflow
	RET
overflow:
	CALL    runtime·abort(SB)
	RET

// PassUnsignedIntegers func(fn unsafe.Pointer, u32 uint32, u64 uint64, u8 uint8, u16 uint16)
TEXT ·PassUnsignedIntegers(SB), $65536-28
	MOVQ    fn+0(FP), AX
	MOVLQZX u32+8(FP), DI
	MOVQ    u64+16(FP), SI
	MOVBQZX u8+24(FP), DX
	MOVWQZX u16+26(FP), CX
	MOVL    $0x26374DD8, R10
	MOVL    R10, 8(SP)
	MOVQ    SP, R12
	LEAQ    65536(SP), SP
	ANDQ    $~15, SP
	CALL    AX
	MOVQ    R12, SP
	MOVL    8(SP), R10
	CMPL    R10, $0x26374DD8
	JNE     overflow
	RET
overflow:
	CALL    runtime·abort(SB)
	RET

// PassFloats func(fn unsafe.Pointer, f32_0 float32, f64_0 float64, f64_1 float64, f32_1 float32)
TEXT ·PassFloats(SB), $65536-36
	MOVQ    fn+0(FP), AX
	MOVSS   f32_0+8(FP), X0
	MOVSD   f64_0+16(FP), X1
	MOVSD   f64_1+24(FP), X2
	MOVSS   f32_1+32(FP), X3
	MOVL    $0x1883CD41, R10
	MOVL    R10, 8(SP)
	MOVQ    SP, R12
	LEAQ    65536(SP), SP
	ANDQ    $~15, SP
	CALL    AX
	MOVQ    R12, SP
	MOVL    8(SP), R10
	CMPL    R10, $0x1883CD41
	JNE     overflow
	RET
overflow:
	CALL    runtime·abort(SB)
	RET

// PassMixedNumbers func(fn unsafe.Pointer, i8 int8, f32 float32, u32 uint32, f64 float64, i64 int64)
TEXT ·PassMixedNumbers(SB), $65536-40
	MOVQ    fn+0(FP), AX
	MOVBQSX i8+8(FP), DI
	MOVSS   f32+12(FP), X0
	MOVLQZX u32+16(FP), SI
	MOVSD   f64+24(FP), X1
	MOVQ    i64+32(FP), DX
	MOVL    $0xBA053D31, R10
	MOVL    R10, 8(SP)
	MOVQ    SP, R12
	LEAQ    65536(SP), SP
	ANDQ    $~15, SP
	CALL    AX
	MOVQ    R12, SP
	MOVL    8(SP), R10
	CMPL    R10, $0xBA053D31
	JNE     overflow
	RET
overflow:
	CALL    runtime·abort(SB)
	RET

// ReturnUInt8 func(fn unsafe.Pointer) uint8
TEXT ·ReturnUInt8(SB), $65536-9
	MOVQ    fn+0(FP), AX
	MOVL    $0xE8B54696, R10
	MOVL    R10, 8(SP)
	MOVQ    SP, R12
	LEAQ    65536(SP), SP
	ANDQ    $~15, SP
	CALL    AX
	MOVQ    R12, SP
	MOVL    8(SP), R10
	CMPL    R10, $0xE8B54696
	JNE     overflow
	MOVB    AX, ret+8(FP)
	RET
overflow:
	CALL    runtime·abort(SB)
	RET

// ReturnInt8 func(fn unsafe.Pointer) int8
TEXT ·ReturnInt8(SB), $65536-9
	MOVQ    fn+0(FP), AX
	MOVL    $0xD0BACDD2, R10
	MOVL    R10, 8(SP)
	MOVQ    SP, R12
	LEAQ    65536(SP), SP
	ANDQ    $~15, SP
	CALL    AX
	MOVQ    R12, SP
	MOVL    8(SP), R10
	CMPL    R10, $0xD0BACDD2
	JNE     overflow
	MOVB    AX, ret+8(FP)
	RET
overflow:
	CALL    runtime·abort(SB)
	RET

// ReturnUInt32 func(fn unsafe.Pointer) uint32
TEXT ·ReturnUInt32(SB), $65536-12
	MOVQ    fn+0(FP), AX
	MOVL    $0x6A240194, R10
	MOVL    R10, 8(SP)
	MOVQ    SP, R12
	LEAQ    65536(SP), SP
	ANDQ    $~15, SP
	CALL    AX
	MOVQ    R12, SP
	MOVL    8(SP), R10
	CMPL    R10, $0x6A240194
	JNE     overflow
	MOVL    AX, ret+8(FP)
	RET
overflow:
	CALL    runtime·abort(SB)
	RET

// ReturnInt32 func(fn unsafe.Pointer) int32
TEXT ·ReturnInt32(SB), $65536-12
	MOVQ    fn+0(FP), AX
	MOVL    $0xE28A748E, R10
	MOVL    R10, 8(SP)
	MOVQ    SP, R12
	LEAQ    65536(SP), SP
	ANDQ    $~15, SP
	CALL    AX
	MOVQ    R12, SP
	MOVL    8(SP), R10
	CMPL    R10, $0xE28A748E
	JNE     overflow
	MOVL    AX, ret+8(FP)
	RET
overflow:
	CALL    runtime·abort(SB)
	RET

// ReturnUInt64 func(fn unsafe.Pointer) uint64
TEXT ·ReturnUInt64(SB), $65536-16
	MOVQ    fn+0(FP), AX
	MOVL    $0xC19949A4, R10
	MOVL    R10, 8(SP)
	MOVQ    SP, R12
	LEAQ    65536(SP), SP
	ANDQ    $~15, SP
	CALL    AX
	MOVQ    R12, SP
	MOVL    8(SP), R10
	CMPL    R10, $0xC19949A4
	JNE     overflow
	MOVQ    AX, ret+8(FP)
	RET
overflow:
	CALL    runtime·abort(SB)
	RET

// ReturnInt64 func(fn unsafe.Pointer) int64
TEXT ·ReturnInt64(SB), $65536-16
	MOVQ    fn+0(FP), AX
	MOVL    $0x6C326E4D, R10
	MOVL    R10, 8(SP)
	MOVQ    SP, R12
	LEAQ    65536(SP), SP
	ANDQ    $~15, SP
	CALL    AX
	MOVQ    R12, SP
	MOVL    8(SP), R10
	CMPL    R10, $0x6C326E4D
	JNE     overflow
	MOVQ    AX, ret+8(FP)
	RET
overflow:
	CALL    runtime·abort(SB)
	RET

// ReturnFloat func(fn unsafe.Pointer) float32
TEXT ·ReturnFloat(SB), $65536-12
	MOVQ    fn+0(FP), AX
	MOVL    $0x92D0DF8F, R10
	MOVL    R10, 8(SP)
	MOVQ    SP, R12
	LEAQ    65536(SP), SP
	ANDQ    $~15, SP
	CALL    AX
	MOVQ    R12, SP
	MOVL    8(SP), R10
	CMPL    R10, $0x92D0DF8F
	JNE     overflow
	MOVSS   X0, ret+8(FP)
	RET
overflow:
	CALL    runtime·abort(SB)
	RET

// ReturnDouble func(fn unsafe.Pointer) float64
TEXT ·ReturnDouble(SB), $65536-16
	MOVQ    fn+0(FP), AX
	MOVL    $0xB9B1D718, R10
	MOVL    R10, 8(SP)
	MOVQ    SP, R12
	LEAQ    65536(SP), SP
	ANDQ    $~15, SP
	CALL    AX
	MOVQ    R12, SP
	MOVL    8(SP), R10
	CMPL    R10, $0xB9B1D718
	JNE     overflow
	MOVSD   X0, ret+8(FP)
	RET
overflow:
	CALL    runtime·abort(SB)
	RET

// PassSmallStructSameIntegers func(fn unsafe.Pointer, s SmallStructSameIntegers)
TEXT ·PassSmallStructSameIntegers(SB), $65536-16
	MOVQ    fn+0(FP), AX
	MOVQ    s_0+8(FP), DI
	MOVL    $0x50EC5FEC, R10
	MOVL    R10, 8(SP)
	MOVQ    SP, R12
	LEAQ    65536(SP), SP
	ANDQ    $~15, SP
	CALL    AX
	MOVQ    R12, SP
	MOVL    8(SP), R10
	CMPL    R10, $0x50EC5FEC
	JNE     overflow
	RET
overflow:
	CALL    runtime·abort(SB)
	RET

// PassSmallStructMixedIntegers func(fn unsafe.Pointer, s SmallStructMixedIntegers)
TEXT ·PassSmallStructMixedIntegers(SB), $65536-26
	MOVQ    fn+0(FP), AX
	MOVQ    s_0+8(FP), DI
	MOVQ    s_1+16(FP), SI
	MOVL    $0xBB5499B4, R10
	MOVL    R10, 8(SP)
	MOVQ    SP, R12
	LEAQ    65536(SP), SP
	ANDQ    $~15, SP
	CALL    AX
	MOVQ    R12, SP
	MOVL    8(SP), R10
	CMPL    R10, $0xBB5499B4
	JNE     overflow
	RET
overflow:
	CALL    runtime·abort(SB)
	RET

// PassSmallStructSameFloats func(fn unsafe.Pointer, s SmallStructSameFloats)
TEXT ·PassSmallStructSameFloats(SB), $65536-28
	MOVQ    fn+0(FP), AX
	MOVSS   s_F32_0+8(FP), X0
	MOVSS   s_F32_1+12(FP), X1
	MOVSS   s_F32_2+16(FP), X2
	MOVL    $0x18D6F455, R10
	MOVL    R10, 8(SP)
	MOVQ    SP, R12
	LEAQ    65536(SP), SP
	ANDQ    $~15, SP
	CALL    AX
	MOVQ    R12, SP
	MOVL    8(SP), R10
	CMPL    R10, $0x18D6F455
	JNE     overflow
	RET
overflow:
	CALL    runtime·abort(SB)
	RET

// PassSmallStructMixedFloats func(fn unsafe.Pointer, s SmallStructMixedFloats)
TEXT ·PassSmallStructMixedFloats(SB), $65536-32
	MOVQ    fn+0(FP), AX
	MOVSS   s_F32_0+8(FP), X0
	MOVSD   s_F64_0+16(FP), X1
	MOVL    $0x4BE45CDC, R10
	MOVL    R10, 8(SP)
	MOVQ    SP, R12
	LEAQ    65536(SP), SP
	ANDQ    $~15, SP
	CALL    AX
	MOVQ    R12, SP
	MOVL    8(SP), R10
	CMPL    R10, $0x4BE45CDC
	JNE     overflow
	RET
overflow:
	CALL    runtime·abort(SB)
	RET

// PassSmallStructMixedNumbers func(fn unsafe.Pointer, s SmallStructMixedNumbers)
TEXT ·PassSmallStructMixedNumbers(SB), $65536-30
	MOVQ    fn+0(FP), AX
	MOVQ    s_0+8(FP), DI
	MOVQ    s_1+16(FP), SI
	MOVL    $0x292B0008, R10
	MOVL    R10, 8(SP)
	MOVQ    SP, R12
	LEAQ    65536(SP), SP
	ANDQ    $~15, SP
	CALL    AX
	MOVQ    R12, SP
	MOVL    8(SP), R10
	CMPL    R10, $0x292B0008
	JNE     overflow
	RET
overflow:
	CALL    runtime·abort(SB)
	RET

// PassSmallStructNested func(fn unsafe.Pointer, s SmallStructOuter)
TEXT ·PassSmallStructNested(SB), $65536-16
	MOVQ    fn+0(FP), AX
	MOVQ    s_0+8(FP), DI
	MOVL    $0xEEF092E7, R10
	MOVL    R10, 8(SP)
	MOVQ    SP, R12
	LEAQ    65536(SP), SP
	ANDQ    $~15, SP
	CALL    AX
	MOVQ    R12, SP
	MOVL    8(SP), R10
	CMPL    R10, $0xEEF092E7
	JNE     overflow
	RET
overflow:
	CALL    runtime·abort(SB)
	RET