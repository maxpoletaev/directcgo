#include "testsuite.h"

#include <stdio.h>
#include <stdint.h>

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