package day09

import (
	"slices"
	"testing"
)

var (
	testInput = []byte(`0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45
`)
	expectedSequences = [][]int{
		{0, 3, 6, 9, 12, 15},
		{1, 3, 6, 10, 15, 21},
		{10, 13, 16, 21, 30, 45},
	}
)

func TestParseInput(t *testing.T) {
	sequences := parseInput(testInput)

	if len(sequences) != len(expectedSequences) {
		t.Fatal("incorrect sequence count")
	}
	for i, s := range sequences {
		es := expectedSequences[i]
		if !slices.Equal(es, s) {
			t.Fatal("incorrect sequence")
		}
	}
}

func TestRun(t *testing.T) {
	result, err := Run(Input)

	if err != nil {
		t.Fatal("unexpected error")
	}
	if result.Part1 != 1666172641 {
		t.Fatal("incorrect part 1")
	}
	//if result.Part2 != 16187743689077 {
	//	t.Fatal("incorrect part 2")
	//}
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result, err := Run(Input)
		_, _ = result, err
	}
}
