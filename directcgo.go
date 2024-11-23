package directcgo

import "unsafe"

//go:noescape
func Call(f unsafe.Pointer, arg unsafe.Pointer, ret unsafe.Pointer)
