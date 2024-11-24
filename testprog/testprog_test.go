package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"testing"
)

// TestProg verifies correctness of directcgo calls under different conditions,
// by running testprog.go, in a separate process:
//  1. Concurrent calls do not interfere with each other.
//  2. GC does not interfere with directcgo calls.
//  3. Stack smashing protection detects overflow at around 1000
//     recursions (each recursion call allocates 1KB stack space).
func TestProg(t *testing.T) {
	var p struct {
		concurrency []int
		depth       []int
	}

	p.concurrency = []int{1, 2, 4}
	p.depth = []int{10, 100, 900, 1000}

	// Permute over all combinations of parameters.
	for _, c := range p.concurrency {
		for _, d := range p.depth {
			t.Run(fmt.Sprintf("concurrency=%d,depth=%d", c, d), func(t *testing.T) {
				cmd := exec.Command(
					"go", "run", "testprog.go",
					"-concurrency", fmt.Sprintf("%d", c),
					"-depth", fmt.Sprintf("%d", d),
					"-silent",
				)

				out, err := cmd.CombinedOutput()

				if d >= 1000 {
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
