package day2502

import (
	"advent/util/tutil"
	_ "embed"
	"testing"
)

//go:embed sample.txt
var sampleInput []byte

const (
	samplePart1 = 1227775554
	samplePart2 = 4174379265
	part1       = 55916882972
	part2       = 76169125915
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
