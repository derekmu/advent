package day07

import (
	"advent/util"
	"bytes"
	_ "embed"
	"time"
)

//go:embed input.txt
var Input []byte
var Problem = util.Problem{Year: "2025", Day: "07", Runner: Run, Input: Input}

func parseInput(input []byte) (lines [][]byte) {
	lines = util.ParseInputLines(input)
	return lines
}

func Run(input []byte) (util.Result, error) {
	start := time.Now()

	lines := parseInput(input)

	parse := time.Now()

	part1 := 0
	part2 := 0

	beamsActive := make(map[int]int, 10)
	beamsNext := make(map[int]int, 10)
	si := bytes.IndexByte(lines[0], 'S')
	beamsActive[si] = 1
	for li := 2; li < len(lines); li += 2 {
		clear(beamsNext)
		for bi, lives := range beamsActive {
			if lines[li][bi] == '^' {
				beamsNext[bi-1] += lives
				beamsNext[bi+1] += lives
				part1++
			} else {
				beamsNext[bi] += lives
			}
		}
		beamsActive, beamsNext = beamsNext, beamsActive
	}
	for _, lives := range beamsActive {
		part2 += lives
	}

	end := time.Now()

	return util.Result{
		Part1:     part1,
		Part2:     part2,
		StartTime: start,
		ParseTime: parse,
		EndTime:   end,
	}, nil
}
