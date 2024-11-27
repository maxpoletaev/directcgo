package testsuite

/*
#include "code.h"
*/
import "C"
import (
	"bytes"
	"unsafe"
)

const (
	outputBufferSize = 65536
)

type Pair struct {
	Key string
	Val string
}

type Pairs []Pair

func (p Pairs) Map() map[string]string {
	m := make(map[string]string)
	for _, pair := range p {
		m[pair.Key] = pair.Val
	}
	return m
}

func getOutput() (pairs Pairs) {
	outBuf := unsafe.Slice(C.GetOutputBuffer(), outputBufferSize)
	var key bytes.Buffer
	var val bytes.Buffer
	var inValue bool

	for _, b := range outBuf {
		if b == ' ' {
			pairs = append(pairs, Pair{
				Key: key.String(),
				Val: val.String(),
			})
			inValue = false
			key.Reset()
			val.Reset()
		} else if b == '=' {
			inValue = true
			continue
		} else if b == 0 {
			break
		} else {
			if inValue {
				val.WriteByte(byte(b))
			} else {
				key.WriteByte(byte(b))
			}
		}
	}

	pairs = append(pairs, Pair{
		Key: key.String(),
		Val: val.String(),
	})

	return pairs
}
