#include "testsuite.h"

char out_buf[MAX_OUTPUT_SIZE];

const char* GetOutputBuffer()
{
    return out_buf;
}

void ResetOutputBuffer()
{
    out_buf[0] = '\0';
}
