package day12

import (
	"testing"
)

var (
	sampleInput = []byte(`Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi
`)
)

func TestRunSample(t *testing.T) {
	result, err := Run(sampleInput)

	if err != nil {
		t.Fatal("unexpected error")
	}
	if result.Part1 != 31 {
		t.Fatal("incorrect part 1")
	}
	if result.Part2 != 29 {
		t.Fatal("incorrect part 2")
	}
}

func TestRun(t *testing.T) {
	result, err := Run(Input)

	if err != nil {
		t.Fatal("unexpected error")
	}
	if result.Part1 != 420 {
		t.Fatal("incorrect part 1")
	}
	if result.Part2 != 414 {
		t.Fatal("incorrect part 2")
	}
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result, err := Run(Input)
		_, _ = result, err
	}
}
