package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"testing"
)

// TestOverflow verifies correctness of directcgo calls under different conditions,
// by running overflow.go, in a separate process. Basically it checks that:
//
//  1. Concurrent calls do not interfere with each other.
//  2. GC triggered at random intervals does not cause issues.
//  3. Stack smashing protection detects overflow at around 64
//     recursions (each recursion call allocates 1KB stack space).
func TestOverflow(t *testing.T) {
	var p struct {
		concurrency []int
		depth       []int
	}

	p.concurrency = []int{1, 2, 4}
	p.depth = []int{5, 10, 30, 50, 60}

	// Permute over all combinations of parameters.
	for _, c := range p.concurrency {
		for _, d := range p.depth {
			t.Run(fmt.Sprintf("concurrency=%d,depth=%d", c, d), func(t *testing.T) {
				cmd := exec.Command(
					"go", "run", "overflow.go",
					"-concurrency", fmt.Sprintf("%d", c),
					"-depth", fmt.Sprintf("%d", d),
					"-silent",
				)

				out, err := cmd.CombinedOutput()

				if d >= 64 {
					if err == nil || !bytes.HasPrefix(out, []byte("SIGSEGV")) {
						t.Fatalf("expected stack smashing protection to trigger segfault\n%s", cmd.String())
					}
				} else {
					if err != nil {
						t.Fatalf("failed to run testprog: %v\n%s", err, out)
					}
				}
			})
		}
	}
}
