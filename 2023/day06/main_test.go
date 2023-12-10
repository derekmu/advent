package day06

import (
	"slices"
	"testing"
)

var (
	sampleInput = []byte(`Time:      7  15   30
Distance:  9  40  200
`)
	expectedTimes    = []int{7, 15, 30}
	expectedRealTime = 71530
	expectedDists    = []int{9, 40, 200}
	expectedRealDist = 940200
)

func TestParseInput(t *testing.T) {
	times, realTime, dists, realDist := parseInput(sampleInput)

	if !slices.Equal(times, expectedTimes) {
		t.Fatal("incorrect times")
	}
	if realTime != expectedRealTime {
		t.Fatal("incorrect realTime")
	}
	if !slices.Equal(dists, expectedDists) {
		t.Fatal("incorrect dists")
	}
	if realDist != expectedRealDist {
		t.Fatal("incorrect realDist")
	}
}

func TestRunSample(t *testing.T) {
	result, err := Run(sampleInput)

	if err != nil {
		t.Fatal("unexpected error")
	}
	if result.Part1 != 288 {
		t.Fatal("incorrect part 1")
	}
	if result.Part2 != 71503 {
		t.Fatal("incorrect part 2")
	}
}

func TestRun(t *testing.T) {
	result, err := Run(Input)

	if err != nil {
		t.Fatal("unexpected error")
	}
	if result.Part1 != 4811940 {
		t.Fatal("incorrect part 1")
	}
	if result.Part2 != 30077773 {
		t.Fatal("incorrect part 2")
	}
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result, err := Run(Input)
		_, _ = result, err
	}
}
