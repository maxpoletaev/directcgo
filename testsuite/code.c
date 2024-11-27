#include "code.h"

#include <stdio.h>
#include <stdint.h>

char out_buf[MAX_OUTPUT_SIZE];

const char* GetOutputBuffer()
{
    return out_buf;
}

void ResetOutputBuffer()
{
    out_buf[0] = '\0';
}

// -----------------------
// Passing primitive types
// -----------------------

void PassIntegers(int32_t i32, int64_t i64, int16_t i16, int8_t i8)
{
    ResetOutputBuffer();
    sprintf(out_buf, "i32=%d i64=%lld i16=%d i8=%d", i32, i64, i16, i8);
}

void PassUnsignedIntegers(uint32_t u32, uint64_t u64, uint8_t u8, uint16_t u16)
{
    ResetOutputBuffer();
    sprintf(out_buf, "u32=%u u64=%llu u8=%u u16=%u", u32, u64, u8, u16);
}

void PassFloats(float f32_0, double f64_0, double f64_1, float f32_1)
{
    ResetOutputBuffer();
    sprintf(out_buf, "f32_0=%f f64_0=%f f64_1=%f f32_1=%f", f32_0, f64_0, f64_1, f32_1);
}

void PassMixedNumbers(int8_t i8, float f32, uint32_t u32, double f64, int64_t i64)
{
    ResetOutputBuffer();
    sprintf(out_buf, "i8=%d f32=%f u32=%u f64=%f i64=%lld", i8, f32, u32, f64, i64);
}

// -------------------------
// Returning primitive types
// -------------------------

uint8_t ReturnUInt8(void)
{
    return 0x12;
}

int8_t ReturnInt8(void)
{
    return -0x12;
}

uint32_t ReturnUInt32(void)
{
    return 0x12345678;
}

int32_t ReturnInt32(void)
{
    return -0x12345678;
}

uint64_t ReturnUInt64(void)
{
    return 0x123456789abcdef0;
}

int64_t ReturnInt64(void)
{
    return -0x123456789abcdef0;
}

float ReturnFloat(void)
{
    return 3.14159f;
}

double ReturnDouble(void)
{
    return 3.14159265358979323846;
}

// ---------------
// Passing structs
// ---------------


void PassSmallStructIntegers(SmallStructIntegers s)
{
    ResetOutputBuffer();
    sprintf(out_buf, "u8=%u i32=%d", s.u8, s.i32);
}

void PassSmallStructFloats(SmallStructFloats s)
{
    ResetOutputBuffer();
    sprintf(out_buf, "f32=%f f64=%f", s.f32, s.f64);
}

void PassSmallStructMixed(SmallStructMixed s)
{
    ResetOutputBuffer();
    sprintf(out_buf, "i32=%d u8=%u f32=%f u16=%u", s.i32, s.u8, s.f32, s.u16);
}

// ------------
// Benchmarking
// ------------

uint32_t AddTwoNumbers(uint32_t a, uint32_t b)
{
    return a + b;
}
