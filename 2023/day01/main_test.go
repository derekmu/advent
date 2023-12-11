package day01

import (
	"testing"
)

var (
	sampleInput = []byte(`two1nine
eightwo6three
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
`)
)

func TestRunSample(t *testing.T) {
	result, err := Run(sampleInput)

	if err != nil {
		t.Fatal("unexpected error")
	}
	if result.Part1 != 275 {
		t.Fatal("incorrect part 1")
	}
	if result.Part2 != 281 {
		t.Fatal("incorrect part 1")
	}
}

func TestRun(t *testing.T) {
	result, err := Run(Input)

	if err != nil {
		t.Fatal("unexpected error")
	}
	if result.Part1 != 56397 {
		t.Fatal("incorrect part 1")
	}
	if result.Part2 != 55701 {
		t.Fatal("incorrect part 2")
	}
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result, err := Run(Input)
		_, _ = result, err
	}
}
