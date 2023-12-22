package tutil

import (
	"advent/util"
	"reflect"
	"testing"
)

type runner func(input []byte) (*util.Result, error)

func RunInput(t testing.TB, run runner, input []byte, p1, p2 any) {
	result, err := run(input)
	if err != nil {
		t.Fatal("unexpected error")
	}
	t.Logf("Part 1: %v, Part 2: %v", result.Part1, result.Part2)
	if !reflect.DeepEqual(result.Part1, p1) {
		t.Fatalf("incorrect part 1, expected %v, got %v", p1, result.Part1)
	}
	if !reflect.DeepEqual(result.Part2, p2) {
		t.Fatalf("incorrect part 2, expected %v, got %v", p2, result.Part2)
	}
}

func BenchInput(b *testing.B, run runner, input []byte, p1, p2 any) {
	for i := 0; i < b.N; i++ {
		RunInput(b, run, input, p1, p2)
	}
}
