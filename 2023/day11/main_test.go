package day11

import (
	"advent/util/tutil"
	_ "embed"
	"testing"
)

//go:embed sample.txt
var sampleInput []byte

const (
	samplePart1 = 374
	samplePart2 = 82000210
	part1       = 10154062
	part2       = 553083047914
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
