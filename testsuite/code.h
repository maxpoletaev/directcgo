#ifndef CODE_H
#define CODE_H

#include <stdint.h>

#define MAX_OUTPUT_SIZE 65536

extern char out_buf[MAX_OUTPUT_SIZE];

const char *GetOutputBuffer();

// -----------------------
// Passing primitive types
// -----------------------

void PassIntegers(int32_t i32, int64_t i64, int16_t i16, int8_t i8);

void PassUnsignedIntegers(uint32_t u32, uint64_t u64, uint8_t u8, uint16_t u16);

void PassFloats(float f32_0, double f64_0, double f64_1, float f32_1);

void PassMixedNumbers(int8_t i8, float f32, uint32_t u32, double f64, int64_t i64);

// -------------------------
// Returning primitive types
// -------------------------

uint8_t ReturnUInt8(void);

int8_t ReturnInt8(void);

uint32_t ReturnUInt32(void);

int32_t ReturnInt32(void);

uint64_t ReturnUInt64(void);

int64_t ReturnInt64(void);

float ReturnFloat(void);

double ReturnDouble(void);

// ---------------
// Passing structs
// ---------------

typedef struct {
    uint8_t u8;
    int32_t i32;
    uint16_t u16;
} SmallStructIntegers;

typedef struct {
    float f32;
    double f64;
} SmallStructFloats;

typedef struct {
    int32_t i32;
    uint8_t u8;
    float f32;
    uint16_t u16;
} SmallStructMixed;

void PassSmallStructIntegers(SmallStructIntegers s);

void PassSmallStructFloats(SmallStructFloats s);

void PassSmallStructMixed(SmallStructMixed s);

// ------------
// Benchmarking
// ------------

uint32_t AddTwoNumbers(uint32_t a, uint32_t b);

#endif
