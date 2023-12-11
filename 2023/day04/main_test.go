package day04

import (
	"testing"
)

var (
	sampleInput = []byte(`Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
`)
	expectedCards = []*card{
		{matches: 4, copies: 1},
		{matches: 2, copies: 1},
		{matches: 2, copies: 1},
		{matches: 1, copies: 1},
		{matches: 0, copies: 1},
		{matches: 0, copies: 1},
	}
)

func TestParseInput(t *testing.T) {
	cards := parseInput(sampleInput)

	if len(cards) != len(expectedCards) {
		t.Fatal("incorrect card count")
	}
	for i, c := range cards {
		if *c != *expectedCards[i] {
			t.Fatal("incorrect card")
		}
	}
}

func TestRunSample(t *testing.T) {
	result, err := Run(sampleInput)

	if err != nil {
		t.Fatal("unexpected error")
	}
	if result.Part1 != 13 {
		t.Fatal("incorrect part 1")
	}
	if result.Part2 != 30 {
		t.Fatal("incorrect part 2")
	}
}

func TestRun(t *testing.T) {
	result, err := Run(Input)

	if err != nil {
		t.Fatal("unexpected error")
	}
	if result.Part1 != 23847 {
		t.Fatal("incorrect part 1")
	}
	if result.Part2 != 8570000 {
		t.Fatal("incorrect part 2")
	}
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result, err := Run(Input)
		_, _ = result, err
	}
}
