package day05

import (
	"slices"
	"testing"
)

var (
	sampleInput = []byte(`seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4
`)
	expectedSeeds = []int{79, 14, 55, 13}
	expectedMaps  = []mapper{
		{
			{52, 50, 97},
			{50, 98, 99},
		},
		{
			{39, 0, 14},
			{0, 15, 51},
			{37, 52, 53},
		},
		{
			{42, 0, 6},
			{57, 7, 10},
			{0, 11, 52},
			{49, 53, 60},
		},
		{
			{88, 18, 24},
			{18, 25, 94},
		},
		{
			{81, 45, 63},
			{68, 64, 76},
			{45, 77, 99},
		},
		{
			{1, 0, 68},
			{0, 69, 69},
		},
		{
			{60, 56, 92},
			{56, 93, 96},
		},
	}
)

func TestParseInput(t *testing.T) {
	seeds, maps := parseInput(sampleInput)

	if !slices.Equal(seeds, expectedSeeds) {
		t.Fatal("incorrect seeds")
	}
	for i := range maps {
		if !slices.Equal(maps[i], expectedMaps[i]) {
			t.Fatal("incorrect maps")
		}
	}
}

func TestRunSample(t *testing.T) {
	result, err := Run(sampleInput)

	if err != nil {
		t.Fatal("unexpected error")
	}
	if result.Part1 != 35 {
		t.Fatal("incorrect part 1")
	}
	if result.Part2 != 46 {
		t.Fatal("incorrect part 2")
	}
}

func TestRun(t *testing.T) {
	result, err := Run(Input)

	if err != nil {
		t.Fatal("unexpected error")
	}
	if result.Part1 != 177942185 {
		t.Fatal("incorrect part 1")
	}
	if result.Part2 != 69841803 {
		t.Fatal("incorrect part 2")
	}
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result, err := Run(Input)
		_, _ = result, err
	}
}
