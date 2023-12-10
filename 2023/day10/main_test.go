package day09

import (
	"testing"
)

var (
	sampleInput = []byte(`7-F7-
.FJ|7
SJLL7
|F--J
LJ.LJ
`)
	expectedStarter = pather{point{2, 0}, none, 0}
)

func TestParseInput(t *testing.T) {
	_, starter := parseInput(sampleInput)

	if starter != expectedStarter {
		t.Fatal("incorrect start location")
	}
}

func TestRunSample1(t *testing.T) {
	result, err := Run(sampleInput)

	if err != nil {
		t.Fatal("unexpected error")
	}
	if result.Part1 != 8 {
		t.Fatal("incorrect part 1")
	}
	if result.Part2 != 1 {
		t.Fatal("incorrect part 2")
	}
}

func TestRun(t *testing.T) {
	result, err := Run(Input)

	if err != nil {
		t.Fatal("unexpected error")
	}
	if result.Part1 != 6951 {
		t.Fatal("incorrect part 1")
	}
	if result.Part2 != 563 {
		t.Fatal("incorrect part 2")
	}
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result, err := Run(Input)
		_, _ = result, err
	}
}
