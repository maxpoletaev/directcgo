package test

import (
	"math/rand"
	"testing"
)

func TestAddTwoNumbersCgo(t *testing.T) {
	for i := 0; i < 1000; i++ {
		var (
			a = rand.Uint32()
			b = rand.Uint32()
		)

		want := AddTwoNumbersNative(a, b)
		got := AddTwoNumbersDirect(a, b)

		if got != want {
			t.Fatalf("want %d, got %d", want, got)
		}
	}
}
