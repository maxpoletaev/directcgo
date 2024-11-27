package codegen

import "testing"

func TestAlign(t *testing.T) {
	tests := []struct {
		offset    int
		alignment int
		want      int
	}{
		// 4-byte alignment
		{9, 4, 12},
		{10, 4, 12},
		{11, 4, 12},

		// 8-byte alignment
		{9, 8, 16},
		{10, 8, 16},
		{11, 8, 16},
		{12, 8, 16},

		// 16-byte alignment
		{9, 16, 16},
		{10, 16, 16},
		{11, 16, 16},
		{20, 16, 32},
		{25, 16, 32},
	}

	for _, tt := range tests {
		if got := align(tt.offset, tt.alignment); got != tt.want {
			t.Errorf("align(%d, %d) = %d; want %d", tt.offset, tt.alignment, got, tt.want)
		}
	}
}
