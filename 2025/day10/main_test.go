package day10

import (
	"advent/util/tutil"
	_ "embed"
	"testing"
)

//go:embed sample.txt
var sampleInput []byte

const (
	samplePart1 = 7
	samplePart2 = -1 // 33
	part1       = 385
	part2       = -1 // didn't solve part 2
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
