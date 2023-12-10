package day09

import (
	"advent/util"
	_ "embed"
	"time"
)

//go:embed input.txt
var Input []byte

func parseInput(input []byte) [][]byte {
	lines := util.ParseInputLines(input)
	return lines
}

func Run(input []byte) (*util.Result, error) {
	start := time.Now()

	_ = parseInput(input)

	parse := time.Now()

	part1 := -1
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
