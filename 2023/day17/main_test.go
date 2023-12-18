package day17

import (
	"advent/util/tutil"
	_ "embed"
	"testing"
)

//go:embed sample.txt
var sampleInput []byte

const (
	samplePart1 = 102
	samplePart2 = -1
	part1       = 1110
	part2       = -1
)

func TestRunSample(t *testing.T) {
	tutil.RunInput(t, Run, sampleInput, samplePart1, samplePart2)
}

func TestRun(t *testing.T) {
	tutil.RunInput(t, Run, Input, part1, part2)
}

func BenchmarkRun(b *testing.B) {
	tutil.BenchInput(b, Run, Input, part1, part2)
}
