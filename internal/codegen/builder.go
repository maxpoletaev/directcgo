package codegen

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
)

type builder struct {
	buf bytes.Buffer
}

func (b *builder) Reset() {
	b.buf.Reset()
}

// NL writes a newline.
func (b *builder) NL() {
	b.buf.WriteByte('\n')
}

// S writes a single line.
func (b *builder) S(format string, args ...any) {
	_, err := fmt.Fprintf(&b.buf, format, args...)
	if err != nil {
		panic(err)
	}
	b.buf.WriteByte('\n')
}

// I writes an assembly instruction with padding.
func (b *builder) I(instr, format string, args ...any) {
	b.buf.WriteByte('\t')

	_, err := fmt.Fprint(&b.buf, instr)
	if err != nil {
		panic(err)
	}

	if format != "" {
		_, err1 := fmt.Fprint(&b.buf, strings.Repeat(" ", 8-len(instr)))
		_, err2 := fmt.Fprintf(&b.buf, format, args...)

		if err = errors.Join(err1, err2); err != nil {
			panic(err)
		}
	}

	b.buf.WriteByte('\n')
}

func (b *builder) Bytes() []byte {
	return b.buf.Bytes()
}
