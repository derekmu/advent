package day03

import (
	"testing"
)

var (
	sampleInput = []byte(`467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
`)
	expectedNumbers = []*number{
		{row: 0, start: 0, end: 2},
		{row: 0, start: 5, end: 7},
		{row: 2, start: 2, end: 3},
		{row: 2, start: 6, end: 8},
		{row: 4, start: 0, end: 2},
		{row: 5, start: 7, end: 8},
		{row: 6, start: 2, end: 4},
		{row: 7, start: 6, end: 8},
		{row: 9, start: 1, end: 3},
		{row: 9, start: 5, end: 7},
	}
	expectedGearMap = map[index][]int{
		{1, 3}: {},
		{4, 3}: {},
		{8, 5}: {},
	}
)

func TestParseInput(t *testing.T) {
	_, numbers, gearMap := parseInput(sampleInput)

	if len(numbers) != len(expectedNumbers) {
		t.Fatal("incorrect numbers")
	}
	for i, n := range numbers {
		if *n != *expectedNumbers[i] {
			t.Fatal("incorrect number")
		}
	}
	if len(gearMap) != len(expectedGearMap) {
		t.Fatal("incorrect gears")
	}
	for key := range gearMap {
		if _, ok := expectedGearMap[key]; !ok {
			t.Fatal("incorrect gears")
		}
	}
}

func TestRunSample(t *testing.T) {
	result, err := Run(sampleInput)

	if err != nil {
		t.Fatal("unexpected error")
	}
	if result.Part1 != 4361 {
		t.Fatal("incorrect part 1")
	}
	if result.Part2 != 467835 {
		t.Fatal("incorrect part 2")
	}
}

func TestRun(t *testing.T) {
	result, err := Run(Input)

	if err != nil {
		t.Fatal("unexpected error")
	}
	if result.Part1 != 527446 {
		t.Fatal("incorrect part 1")
	}
	if result.Part2 != 73201705 {
		t.Fatal("incorrect part 2")
	}
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result, err := Run(Input)
		_, _ = result, err
	}
}
