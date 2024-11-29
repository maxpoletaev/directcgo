#include "testsuite.h"
#include <stdio.h>
#include <stdint.h>

// ---------------
// Passing structs
// ---------------

void PassSmallStructSameIntegers(SmallStructSameIntegers s)
{
    ResetOutputBuffer();
    sprintf(out_buf, "u32_0=%u u32_1=%u", s.u32_0, s.u32_1);
}

void PassSmallStructMixedIntegers(SmallStructMixedIntegers s)
{
    ResetOutputBuffer();
    sprintf(out_buf, "u8=%u i32=%d u16=%u", s.u8, s.i32, s.u16);
}

void PassSmallStructSameFloats(SmallStructSameFloats s)
{
    ResetOutputBuffer();
    sprintf(out_buf, "f32_0=%f f32_1=%f f32_2=%f", s.f32_0, s.f32_1, s.f32_2);
}

void PassSmallStructMixedFloats(SmallStructMixedFloats s)
{
    ResetOutputBuffer();
    sprintf(out_buf, "f32_0=%f f64_0=%f", s.f32_0, s.f64_0);
}

void PassSmallStructMixedNumbers(SmallStructMixedNumbers s)
{
    ResetOutputBuffer();
    sprintf(out_buf, "i32=%d u8=%u f32=%f u16=%u", s.i32, s.u8, s.f32, s.u16);
}

void PassSmallStructNested(SmallStructOuter s)
{
    ResetOutputBuffer();
    sprintf(out_buf, "inner_0.u32=%u inner_1.u32=%u", s.inner_0.u32, s.inner_1.u32);
}

void PassSmallStructWithArray(SmallStructWithArray s)
{
    ResetOutputBuffer();
    sprintf(out_buf, "u8=%u arr=[%u,%u,%u] f64=%f", s.u8, s.arr[0], s.arr[1], s.arr[2], s.f64);
}
