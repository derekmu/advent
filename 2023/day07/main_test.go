package day07

import (
	"advent/util/tutil"
	_ "embed"
	"testing"
)

//go:embed sample.txt
var sampleInput []byte

const (
	samplePart1 = 6440
	samplePart2 = 5905
	part1       = 253866470
	part2       = 254494947
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
