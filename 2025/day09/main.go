package day09

import (
	"advent/util"
	_ "embed"
	"time"
)

//go:embed input.txt
var Input []byte
var Problem = util.Problem{Year: "2025", Day: "09", Runner: Run, Input: Input}

type point struct {
	x, y int
}

func parseInput(input []byte) (points []point) {
	lines := util.ParseInputLines(input)
	points = make([]point, len(lines))
	for i, line := range lines {
		parts := util.ParseInputDelimiter(line, []byte(","))
		points[i] = point{
			x: util.Btoi(parts[0]),
			y: util.Btoi(parts[1]),
		}
	}
	return points
}

func Run(input []byte) (util.Result, error) {
	start := time.Now()

	points := parseInput(input)

	parse := time.Now()

	part1 := -1
	part2 := -1

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			dx := util.Abs(points[i].x-points[j].x) + 1
			dy := util.Abs(points[i].y-points[j].y) + 1
			if dx*dy > part1 {
				part1 = dx * dy
			}
		}
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
