package day03

import (
	"advent/util"
	_ "embed"
	"time"
)

//go:embed input.txt
var Input []byte
var Problem = util.Problem{Year: "2025", Day: "03", Runner: Run, Input: Input}

func parseInput(input []byte) (lines [][]byte) {
	input = util.CopyInput(input)
	lines = util.ParseInputLines(input)
	for _, line := range lines {
		for i := range line {
			line[i] -= '0'
		}
	}
	return lines
}

func Run(input []byte) (*util.Result, error) {
	start := time.Now()

	lines := parseInput(input)

	parse := time.Now()

	part1 := 0
	part2 := 0

	for _, line := range lines {
		li := 0
		ri := len(line) - 1
		for tli := li + 1; tli < ri; tli++ {
			if line[tli] > line[li] {
				li = tli
			}
		}
		for tri := ri - 1; tri > li; tri-- {
			if line[tri] > line[ri] {
				ri = tri
			}
		}
		part1 += int(line[li])*10 + int(line[ri])

		li = 0
		for digit := 11; digit >= 0; digit-- {
			for tli := li + 1; tli < len(line)-digit; tli++ {
				if line[tli] > line[li] {
					li = tli
				}
			}
			part2 += util.Pow(10, digit) * int(line[li])
			li++
		}
	}

	end := time.Now()

	return &util.Result{
		Part1:     part1,
		Part2:     part2,
		StartTime: start,
		ParseTime: parse,
		EndTime:   end,
	}, nil
}
