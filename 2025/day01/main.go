package day2501

import (
	"advent/util"
	_ "embed"
	"time"
)

//go:embed input.txt
var Input []byte
var Problem = util.Problem{Year: "2025", Day: "01", Runner: Run, Input: Input}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func parseInput(input []byte) (spins []int) {
	lines := util.ParseInputLines(input)
	spins = make([]int, len(lines))
	for i, line := range lines {
		c := util.Btoi(line[1:])
		if line[0] == 'L' {
			c = -c
		}
		spins[i] = c
	}
	return spins
}

func Run(input []byte) (*util.Result, error) {
	start := time.Now()

	spins := parseInput(input)

	parse := time.Now()

	part1 := 0
	part2 := 0

	current := 50
	for _, spin := range spins {
		before := current
		current = current + spin
		if current <= 0 {
			if before == 0 {
				part2--
			}
			part2 += 1 + abs(current)/100
		} else if current > 99 {
			part2 += current / 100
		}
		current = current % 100
		if current < 0 {
			current += 100
		}
		if current == 0 {
			part1++
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
