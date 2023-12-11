package day02

import (
	"testing"
)

var (
	sampleInput = []byte(`Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
`)
	expectedGames = []*game{
		{
			gameId: 1,
			sets: []*set{
				{red: 4, green: 0, blue: 3},
				{red: 1, green: 2, blue: 6},
				{red: 0, green: 2, blue: 0},
			},
		},
		{
			gameId: 2,
			sets: []*set{
				{red: 0, green: 2, blue: 1},
				{red: 1, green: 3, blue: 4},
				{red: 0, green: 1, blue: 1},
			},
		},
		{
			gameId: 3,
			sets: []*set{
				{red: 20, green: 8, blue: 6},
				{red: 4, green: 13, blue: 5},
				{red: 1, green: 5, blue: 0},
			},
		},
		{
			gameId: 4,
			sets: []*set{
				{red: 3, green: 1, blue: 6},
				{red: 6, green: 3, blue: 0},
				{red: 14, green: 3, blue: 15},
			},
		},
		{
			gameId: 5,
			sets: []*set{
				{red: 6, green: 3, blue: 1},
				{red: 1, green: 2, blue: 2},
			},
		},
	}
)

func TestParseInput(t *testing.T) {
	games, err := parseInput(sampleInput)

	if err != nil {
		t.Fatal("unexpected error")
	}
	if len(games) != len(expectedGames) {
		t.Fatal("incorrect games")
	}
	for i, g := range games {
		eg := expectedGames[i]
		if g.gameId != eg.gameId {
			t.Fatal("incorrect gameId")
		}
		if len(g.sets) != len(eg.sets) {
			t.Fatal("incorrect sets")
		}
		for si, s := range g.sets {
			if *s != *(eg.sets[si]) {
				t.Fatal("incorrect set")
			}
		}
	}
}

func TestRunSample(t *testing.T) {
	result, err := Run(sampleInput)

	if err != nil {
		t.Fatal("unexpected error")
	}
	if result.Part1 != 8 {
		t.Fatal("incorrect part 1")
	}
	if result.Part2 != 2286 {
		t.Fatal("incorrect part 2")
	}
}

func TestRun(t *testing.T) {
	result, err := Run(Input)

	if err != nil {
		t.Fatal("unexpected error")
	}
	if result.Part1 != 2879 {
		t.Fatal("incorrect part 1")
	}
	if result.Part2 != 65122 {
		t.Fatal("incorrect part 2")
	}
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result, err := Run(Input)
		_, _ = result, err
	}
}
