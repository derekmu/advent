package day15

import (
	"advent/util"
	"bytes"
	_ "embed"
	"time"
)

//go:embed input.txt
var Input []byte

func parseInput(input []byte) (lines [][]byte) {
	input = input[:len(input)-1] // strip new line
	return bytes.Split(input, []byte(","))
}

func Run(input []byte) (*util.Result, error) {
	start := time.Now()

	lines := parseInput(input)

	parse := time.Now()

	part1 := 0
	for _, line := range lines {
		cv := 0
		for _, c := range line {
			cv += int(c)
			cv *= 17
			cv = cv % 256
		}
		part1 += cv
	}
	part2 := -1

	end := time.Now()

	return &util.Result{
		Part1:     part1,
		Part2:     part2,
		StartTime: start,
		ParseTime: parse,
		EndTime:   end,
	}, nil
}
