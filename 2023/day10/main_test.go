package day09

import (
	"slices"
	"testing"
)

var (
	sampleInput = []byte(``)
	expectedXXX = [][]byte{}
)

func TestParseInput(t *testing.T) {
	xxx := parseInput(sampleInput)

	if len(xxx) != len(expectedXXX) {
		t.Fatal("incorrect sequence count")
	}
	for i, s := range xxx {
		es := expectedXXX[i]
		if !slices.Equal(es, s) {
			t.Fatal("incorrect sequence")
		}
	}
}

func TestRunSample(t *testing.T) {
	result, err := Run(sampleInput)

	if err != nil {
		t.Fatal("unexpected error")
	}
	if result.Part1 != -1 {
		t.Fatal("incorrect part 1")
	}
	if result.Part2 != -1 {
		t.Fatal("incorrect part 2")
	}
}

func TestRun(t *testing.T) {
	result, err := Run(Input)

	if err != nil {
		t.Fatal("unexpected error")
	}
	if result.Part1 != -1 {
		t.Fatal("incorrect part 1")
	}
	if result.Part2 != -1 {
		t.Fatal("incorrect part 2")
	}
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result, err := Run(Input)
		_, _ = result, err
	}
}
