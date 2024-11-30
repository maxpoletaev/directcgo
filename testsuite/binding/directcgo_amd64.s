//go:build amd64 && !windows

// Code generated by directcgo. DO NOT EDIT.
// directcgo -arch=amd64,arm64 ./testsuite/binding

#include "go_asm.h"
#include "textflag.h"
#include "funcdata.h"

// PassIntegers func(fn unsafe.Pointer, i32 int32, i64 int64, i16 int16, i8 int8)
TEXT ·PassIntegers(SB), $65536-27
	MOVQ    SP, R12
	LEAQ    65536(SP), R11
	ANDQ    $~15, R11
	SUBQ    $0, R11
	MOVQ    fn+0(FP), AX
	MOVLQSX i32+8(FP), DI
	MOVQ    i64+16(FP), SI
	MOVWQSX i16+24(FP), DX
	MOVBQSX i8+26(FP), CX
	MOVL    $0xF51A4219, R10
	MOVL    R10, 8(R12)
	MOVQ    R11, SP
	CALL    AX
	MOVQ    R12, SP
	MOVL    8(SP), R10
	CMPL    R10, $0xF51A4219
	JNE     overflow
	RET
overflow:
	CALL    runtime·abort(SB)
	RET

// PassUnsignedIntegers func(fn unsafe.Pointer, u32 uint32, u64 uint64, u8 uint8, u16 uint16)
TEXT ·PassUnsignedIntegers(SB), $65536-28
	MOVQ    SP, R12
	LEAQ    65536(SP), R11
	ANDQ    $~15, R11
	SUBQ    $0, R11
	MOVQ    fn+0(FP), AX
	MOVLQZX u32+8(FP), DI
	MOVQ    u64+16(FP), SI
	MOVBQZX u8+24(FP), DX
	MOVWQZX u16+26(FP), CX
	MOVL    $0x4A36BCE, R10
	MOVL    R10, 8(R12)
	MOVQ    R11, SP
	CALL    AX
	MOVQ    R12, SP
	MOVL    8(SP), R10
	CMPL    R10, $0x4A36BCE
	JNE     overflow
	RET
overflow:
	CALL    runtime·abort(SB)
	RET

// PassFloats func(fn unsafe.Pointer, f32_0 float32, f64_0 float64, f64_1 float64, f32_1 float32)
TEXT ·PassFloats(SB), $65536-36
	MOVQ    SP, R12
	LEAQ    65536(SP), R11
	ANDQ    $~15, R11
	SUBQ    $0, R11
	MOVQ    fn+0(FP), AX
	MOVSS   f32_0+8(FP), X0
	MOVSD   f64_0+16(FP), X1
	MOVSD   f64_1+24(FP), X2
	MOVSS   f32_1+32(FP), X3
	MOVL    $0xFB966F8A, R10
	MOVL    R10, 8(R12)
	MOVQ    R11, SP
	CALL    AX
	MOVQ    R12, SP
	MOVL    8(SP), R10
	CMPL    R10, $0xFB966F8A
	JNE     overflow
	RET
overflow:
	CALL    runtime·abort(SB)
	RET

// PassMixedNumbers func(fn unsafe.Pointer, i8 int8, f32 float32, u32 uint32, f64 float64, i64 int64)
TEXT ·PassMixedNumbers(SB), $65536-40
	MOVQ    SP, R12
	LEAQ    65536(SP), R11
	ANDQ    $~15, R11
	SUBQ    $0, R11
	MOVQ    fn+0(FP), AX
	MOVBQSX i8+8(FP), DI
	MOVSS   f32+12(FP), X0
	MOVLQZX u32+16(FP), SI
	MOVSD   f64+24(FP), X1
	MOVQ    i64+32(FP), DX
	MOVL    $0x61A92597, R10
	MOVL    R10, 8(R12)
	MOVQ    R11, SP
	CALL    AX
	MOVQ    R12, SP
	MOVL    8(SP), R10
	CMPL    R10, $0x61A92597
	JNE     overflow
	RET
overflow:
	CALL    runtime·abort(SB)
	RET

// ReturnUInt8 func(fn unsafe.Pointer) uint8
TEXT ·ReturnUInt8(SB), $65536-9
	MOVQ    SP, R12
	LEAQ    65536(SP), R11
	ANDQ    $~15, R11
	SUBQ    $0, R11
	MOVQ    fn+0(FP), AX
	MOVL    $0x8A5FE07A, R10
	MOVL    R10, 8(R12)
	MOVQ    R11, SP
	CALL    AX
	MOVQ    R12, SP
	MOVL    8(SP), R10
	CMPL    R10, $0x8A5FE07A
	JNE     overflow
	MOVB    AX, ret+8(FP)
	RET
overflow:
	CALL    runtime·abort(SB)
	RET

// ReturnInt8 func(fn unsafe.Pointer) int8
TEXT ·ReturnInt8(SB), $65536-9
	MOVQ    SP, R12
	LEAQ    65536(SP), R11
	ANDQ    $~15, R11
	SUBQ    $0, R11
	MOVQ    fn+0(FP), AX
	MOVL    $0xC188B80B, R10
	MOVL    R10, 8(R12)
	MOVQ    R11, SP
	CALL    AX
	MOVQ    R12, SP
	MOVL    8(SP), R10
	CMPL    R10, $0xC188B80B
	JNE     overflow
	MOVB    AX, ret+8(FP)
	RET
overflow:
	CALL    runtime·abort(SB)
	RET

// ReturnUInt32 func(fn unsafe.Pointer) uint32
TEXT ·ReturnUInt32(SB), $65536-12
	MOVQ    SP, R12
	LEAQ    65536(SP), R11
	ANDQ    $~15, R11
	SUBQ    $0, R11
	MOVQ    fn+0(FP), AX
	MOVL    $0x2414C9BE, R10
	MOVL    R10, 8(R12)
	MOVQ    R11, SP
	CALL    AX
	MOVQ    R12, SP
	MOVL    8(SP), R10
	CMPL    R10, $0x2414C9BE
	JNE     overflow
	MOVL    AX, ret+8(FP)
	RET
overflow:
	CALL    runtime·abort(SB)
	RET

// ReturnInt32 func(fn unsafe.Pointer) int32
TEXT ·ReturnInt32(SB), $65536-12
	MOVQ    SP, R12
	LEAQ    65536(SP), R11
	ANDQ    $~15, R11
	SUBQ    $0, R11
	MOVQ    fn+0(FP), AX
	MOVL    $0x7EE47F, R10
	MOVL    R10, 8(R12)
	MOVQ    R11, SP
	CALL    AX
	MOVQ    R12, SP
	MOVL    8(SP), R10
	CMPL    R10, $0x7EE47F
	JNE     overflow
	MOVL    AX, ret+8(FP)
	RET
overflow:
	CALL    runtime·abort(SB)
	RET

// ReturnUInt64 func(fn unsafe.Pointer) uint64
TEXT ·ReturnUInt64(SB), $65536-16
	MOVQ    SP, R12
	LEAQ    65536(SP), R11
	ANDQ    $~15, R11
	SUBQ    $0, R11
	MOVQ    fn+0(FP), AX
	MOVL    $0x87B7B8F8, R10
	MOVL    R10, 8(R12)
	MOVQ    R11, SP
	CALL    AX
	MOVQ    R12, SP
	MOVL    8(SP), R10
	CMPL    R10, $0x87B7B8F8
	JNE     overflow
	MOVQ    AX, ret+8(FP)
	RET
overflow:
	CALL    runtime·abort(SB)
	RET

// ReturnInt64 func(fn unsafe.Pointer) int64
TEXT ·ReturnInt64(SB), $65536-16
	MOVQ    SP, R12
	LEAQ    65536(SP), R11
	ANDQ    $~15, R11
	SUBQ    $0, R11
	MOVQ    fn+0(FP), AX
	MOVL    $0xA760BDE5, R10
	MOVL    R10, 8(R12)
	MOVQ    R11, SP
	CALL    AX
	MOVQ    R12, SP
	MOVL    8(SP), R10
	CMPL    R10, $0xA760BDE5
	JNE     overflow
	MOVQ    AX, ret+8(FP)
	RET
overflow:
	CALL    runtime·abort(SB)
	RET

// ReturnFloat func(fn unsafe.Pointer) float32
TEXT ·ReturnFloat(SB), $65536-12
	MOVQ    SP, R12
	LEAQ    65536(SP), R11
	ANDQ    $~15, R11
	SUBQ    $0, R11
	MOVQ    fn+0(FP), AX
	MOVL    $0x89131C30, R10
	MOVL    R10, 8(R12)
	MOVQ    R11, SP
	CALL    AX
	MOVQ    R12, SP
	MOVL    8(SP), R10
	CMPL    R10, $0x89131C30
	JNE     overflow
	MOVSS   X0, ret+8(FP)
	RET
overflow:
	CALL    runtime·abort(SB)
	RET

// ReturnDouble func(fn unsafe.Pointer) float64
TEXT ·ReturnDouble(SB), $65536-16
	MOVQ    SP, R12
	LEAQ    65536(SP), R11
	ANDQ    $~15, R11
	SUBQ    $0, R11
	MOVQ    fn+0(FP), AX
	MOVL    $0x47EA2CF, R10
	MOVL    R10, 8(R12)
	MOVQ    R11, SP
	CALL    AX
	MOVQ    R12, SP
	MOVL    8(SP), R10
	CMPL    R10, $0x47EA2CF
	JNE     overflow
	MOVSD   X0, ret+8(FP)
	RET
overflow:
	CALL    runtime·abort(SB)
	RET

// PassSmallStructSameIntegers func(fn unsafe.Pointer, s SmallStructSameIntegers)
TEXT ·PassSmallStructSameIntegers(SB), $65536-16
	MOVQ    SP, R12
	LEAQ    65536(SP), R11
	ANDQ    $~15, R11
	SUBQ    $0, R11
	MOVQ    fn+0(FP), AX
	MOVQ    s+8(FP), DI
	MOVL    $0x8B8B10C2, R10
	MOVL    R10, 8(R12)
	MOVQ    R11, SP
	CALL    AX
	MOVQ    R12, SP
	MOVL    8(SP), R10
	CMPL    R10, $0x8B8B10C2
	JNE     overflow
	RET
overflow:
	CALL    runtime·abort(SB)
	RET

// PassSmallStructMixedIntegers func(fn unsafe.Pointer, s SmallStructMixedIntegers)
TEXT ·PassSmallStructMixedIntegers(SB), $65536-26
	MOVQ    SP, R12
	LEAQ    65536(SP), R11
	ANDQ    $~15, R11
	SUBQ    $0, R11
	MOVQ    fn+0(FP), AX
	MOVQ    s+8(FP), DI
	MOVQ    s+16(FP), SI
	MOVL    $0x879EB4D2, R10
	MOVL    R10, 8(R12)
	MOVQ    R11, SP
	CALL    AX
	MOVQ    R12, SP
	MOVL    8(SP), R10
	CMPL    R10, $0x879EB4D2
	JNE     overflow
	RET
overflow:
	CALL    runtime·abort(SB)
	RET

// PassSmallStructSameFloats func(fn unsafe.Pointer, s SmallStructSameFloats)
TEXT ·PassSmallStructSameFloats(SB), $65536-28
	MOVQ    SP, R12
	LEAQ    65536(SP), R11
	ANDQ    $~15, R11
	SUBQ    $0, R11
	MOVQ    fn+0(FP), AX
	MOVQ    s+8(FP), X0
	MOVQ    s+16(FP), X1
	MOVL    $0xA8292C92, R10
	MOVL    R10, 8(R12)
	MOVQ    R11, SP
	CALL    AX
	MOVQ    R12, SP
	MOVL    8(SP), R10
	CMPL    R10, $0xA8292C92
	JNE     overflow
	RET
overflow:
	CALL    runtime·abort(SB)
	RET

// PassSmallStructMixedFloats func(fn unsafe.Pointer, s SmallStructMixedFloats)
TEXT ·PassSmallStructMixedFloats(SB), $65536-32
	MOVQ    SP, R12
	LEAQ    65536(SP), R11
	ANDQ    $~15, R11
	SUBQ    $0, R11
	MOVQ    fn+0(FP), AX
	MOVQ    s+8(FP), X0
	MOVQ    s+16(FP), X1
	MOVL    $0x80DD353C, R10
	MOVL    R10, 8(R12)
	MOVQ    R11, SP
	CALL    AX
	MOVQ    R12, SP
	MOVL    8(SP), R10
	CMPL    R10, $0x80DD353C
	JNE     overflow
	RET
overflow:
	CALL    runtime·abort(SB)
	RET

// PassSmallStructMixedNumbers func(fn unsafe.Pointer, s SmallStructMixedNumbers)
TEXT ·PassSmallStructMixedNumbers(SB), $65536-30
	MOVQ    SP, R12
	LEAQ    65536(SP), R11
	ANDQ    $~15, R11
	SUBQ    $0, R11
	MOVQ    fn+0(FP), AX
	MOVQ    s+8(FP), DI
	MOVQ    s+16(FP), SI
	MOVL    $0xF126DC84, R10
	MOVL    R10, 8(R12)
	MOVQ    R11, SP
	CALL    AX
	MOVQ    R12, SP
	MOVL    8(SP), R10
	CMPL    R10, $0xF126DC84
	JNE     overflow
	RET
overflow:
	CALL    runtime·abort(SB)
	RET

// PassSmallStructNested func(fn unsafe.Pointer, s SmallStructOuter)
TEXT ·PassSmallStructNested(SB), $65536-16
	MOVQ    SP, R12
	LEAQ    65536(SP), R11
	ANDQ    $~15, R11
	SUBQ    $0, R11
	MOVQ    fn+0(FP), AX
	MOVQ    s+8(FP), DI
	MOVL    $0xDA593CF1, R10
	MOVL    R10, 8(R12)
	MOVQ    R11, SP
	CALL    AX
	MOVQ    R12, SP
	MOVL    8(SP), R10
	CMPL    R10, $0xDA593CF1
	JNE     overflow
	RET
overflow:
	CALL    runtime·abort(SB)
	RET

// PassSmallStructWithArray func(fn unsafe.Pointer, s SmallStructWithArray)
TEXT ·PassSmallStructWithArray(SB), $65536-32
	MOVQ    SP, R12
	LEAQ    65536(SP), R11
	ANDQ    $~15, R11
	SUBQ    $0, R11
	MOVQ    fn+0(FP), AX
	MOVQ    s+8(FP), DI
	MOVQ    s+16(FP), X0
	MOVL    $0xEAB8C376, R10
	MOVL    R10, 8(R12)
	MOVQ    R11, SP
	CALL    AX
	MOVQ    R12, SP
	MOVL    8(SP), R10
	CMPL    R10, $0xEAB8C376
	JNE     overflow
	RET
overflow:
	CALL    runtime·abort(SB)
	RET

// PassLargeStruct func(fn unsafe.Pointer, s LargeStruct)
TEXT ·PassLargeStruct(SB), $65536-48
	MOVQ    SP, R12
	LEAQ    65536(SP), R11
	ANDQ    $~15, R11
	SUBQ    $0, R11
	MOVQ    fn+0(FP), AX
	MOVQ    s+8(FP), R10
	MOVQ    R10, 0(R11)
	MOVQ    s+16(FP), R10
	MOVQ    R10, 8(R11)
	MOVQ    s+24(FP), R10
	MOVQ    R10, 16(R11)
	MOVQ    s+32(FP), R10
	MOVQ    R10, 24(R11)
	MOVQ    s+40(FP), R10
	MOVQ    R10, 32(R11)
	MOVL    $0x2ADF3FE8, R10
	MOVL    R10, 8(R12)
	MOVQ    R11, SP
	CALL    AX
	MOVQ    R12, SP
	MOVL    8(SP), R10
	CMPL    R10, $0x2ADF3FE8
	JNE     overflow
	RET
overflow:
	CALL    runtime·abort(SB)
	RET
