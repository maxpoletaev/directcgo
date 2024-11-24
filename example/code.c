#include "code.h"
#include <stdio.h>
#include <stdint.h>

void PrintHello(char *name)
{
    printf("Hello, %s!\n", name);
}

uint32_t AddTwoNumbers(uint32_t a, uint32_t b)
{
    return a + b;
}
