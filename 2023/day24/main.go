package day24

import (
	"advent/util"
	"bytes"
	_ "embed"
	"time"
)

//go:embed input.txt
var Input []byte
var Problem = util.Problem{Year: "2023", Day: "XX", Runner: Run, Input: Input}

type stone struct {
	x, y, z    float64
	dx, dy, dz float64
}

func parseInput(input []byte) (stones []stone, minint, maxint float64) {
	lines := util.ParseInputLines(input)
	line1 := lines[0]
	minints, maxints, _ := bytes.Cut(line1, []byte(", "))
	minint = float64(util.Btoi(minints))
	maxint = float64(util.Btoi(maxints))
	lines = lines[1:]
	stones = make([]stone, len(lines))
	var x, y, z, dx, dy, dz []byte
	for i, line := range lines {
		x, line, _ = bytes.Cut(line, []byte(", "))
		y, line, _ = bytes.Cut(line, []byte(", "))
		z, line, _ = bytes.Cut(line, []byte(" @ "))
		dx, line, _ = bytes.Cut(line, []byte(", "))
		dy, dz, _ = bytes.Cut(line, []byte(", "))
		s := stone{
			x:  float64(util.Btoi(x)),
			y:  float64(util.Btoi(y)),
			z:  float64(util.Btoi(z)),
			dx: float64(util.Btoi(dx)),
			dy: float64(util.Btoi(dy)),
			dz: float64(util.Btoi(dz)),
		}
		stones[i] = s
	}
	return stones, minint, maxint
}

func Run(input []byte) (*util.Result, error) {
	start := time.Now()

	stones, mini, maxi := parseInput(input)

	parse := time.Now()

	part1 := 0
	for i, s1 := range stones {
		for j := i + 1; j < len(stones); j++ {
			s2 := stones[j]
			det := s1.dx*s2.dy - s1.dy*s2.dx

			t := ((s2.x-s1.x)*s2.dy - (s2.y-s1.y)*s2.dx) / det
			u := ((s2.x-s1.x)*s1.dy - (s2.y-s1.y)*s1.dx) / det

			if t >= 0 && u >= 0 {
				xi := s1.x + t*s1.dx
				yi := s1.y + t*s1.dy
				if xi >= mini && xi <= maxi && yi >= mini && yi <= maxi {
					part1++
				}
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
