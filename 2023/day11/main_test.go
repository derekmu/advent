package day09

import (
	"testing"
)

var sampleInput = []byte(`...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....
`)

func TestParseInput(t *testing.T) {
	parseInput(sampleInput)
}

func TestRunSample1(t *testing.T) {
	result, err := Run(sampleInput)
	if err != nil {
		t.Fatal("unexpected error")
	}
	if result.Part1 != 374 {
		t.Fatalf("incorrect part 1, expected 374, got %d", result.Part1)
	}
	if result.Part2 != 82000210 {
		t.Fatal("incorrect part 2")
	}
}

func TestRun(t *testing.T) {
	result, err := Run(Input)
	if err != nil {
		t.Fatal("unexpected error")
	}
	if result.Part1 != 10154062 {
		t.Fatal("incorrect part 1")
	}
	if result.Part2 != 553083047914 {
		t.Fatal("incorrect part 2")
	}
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result, err := Run(Input)
		_, _ = result, err
	}
}
