package day14

import (
	"advent/util"
	_ "embed"
	"time"
)

//go:embed input.txt
var Input []byte

func parseInput(input []byte) (lines [][]byte) {
	lines = util.ParseInputLines(input)
	return lines
}

func Run(input []byte) (*util.Result, error) {
	start := time.Now()

	lines := parseInput(input)

	parse := time.Now()

	cols := len(lines[0])

	part1 := 0
	for c := 0; c < cols; c++ {
		r1 := 0
		for r, line := range lines {
			switch line[c] {
			case 'O':
				part1 += cols - r1
				r1++
			case '#':
				r1 = r + 1
			case '.':
			default:
				panic("invalid character")
			}
		}
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
