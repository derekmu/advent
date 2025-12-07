package template

import (
	"advent/util"
	_ "embed"
	"time"
)

//go:embed input.txt
var Input []byte
var Problem = util.Problem{Year: "2025", Day: "XX", Runner: Run, Input: Input}

func parseInput(input []byte) (lines [][]byte) {
	lines = util.ParseInputLines(input)
	return lines
}

func Run(input []byte) (util.Result, error) {
	start := time.Now()

	lines := parseInput(input)
	_ = lines

	parse := time.Now()

	part1 := -1
	part2 := -1

	end := time.Now()

	return util.Result{
		Part1:     part1,
		Part2:     part2,
		StartTime: start,
		ParseTime: parse,
		EndTime:   end,
	}, nil
}
