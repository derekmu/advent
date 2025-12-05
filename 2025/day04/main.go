package day04

import (
	"advent/util"
	_ "embed"
	"time"
)

//go:embed input.txt
var Input []byte
var Problem = util.Problem{Year: "2025", Day: "04", Runner: Run, Input: Input}

func parseInput(input []byte) (lines [][]byte) {
	input = util.CopyInput(input)
	for i := range input {
		if input[i] == '.' {
			input[i] = 0
		} else if input[i] == '@' {
			input[i] = 1
		}
	}
	lines = util.ParseInputLines(input)
	return lines
}

func Run(input []byte) (*util.Result, error) {
	start := time.Now()

	lines := parseInput(input)

	parse := time.Now()

	part1 := 0
	part2 := 0

	width := len(lines[0])
	height := len(lines)

	for y := range lines {
		for x := range lines[y] {
			if lines[y][x] == 0 {
				continue
			}
			c := byte(1)
			for d := range 8 {
				if d > 3 {
					d++
				}
				tx := x + d%3 - 1
				if tx < 0 || tx >= width {
					continue
				}
				ty := y + d/3 - 1
				if ty < 0 || ty >= height {
					continue
				}
				if lines[ty][tx] > 0 {
					c++
				}
			}
			lines[y][x] = c
			if c <= 4 {
				part1++
			}
		}
	}

	x, y := 0, 0
	for y < height {
		part2 += checkRemove(lines, x, y)
		x++
		if x >= width {
			x = 0
			y++
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

func checkRemove(lines [][]byte, x, y int) (removed int) {
	v := lines[y][x]
	if v > 0 && v <= 4 {
		lines[y][x] = 0
		removed++
		for d := range 8 {
			if d > 3 {
				d++
			}
			tx := x + d%3 - 1
			if tx < 0 || tx >= len(lines[0]) {
				continue
			}
			ty := y + d/3 - 1
			if ty < 0 || ty >= len(lines) {
				continue
			}
			lines[ty][tx]--
			removed += checkRemove(lines, tx, ty)
		}
	}
	return removed
}
