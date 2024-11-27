package asm

import "unsafe"

//go:noescape
func AddTwoNumbers(fn unsafe.Pointer, a, b uint32) uint32
