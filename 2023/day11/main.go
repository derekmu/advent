package day11

import (
	"advent/util"
	_ "embed"
	"time"
)

var Problem = util.Problem{Year: "2023", Day: "11", Runner: Run, Input: Input}

type galaxy struct {
	row1 int
	col1 int
	row2 int
	col2 int
}

//go:embed input.txt
var Input []byte

func parseInput(input []byte) (galaxies []galaxy) {
	lines := util.ParseInputLines(input)
	galaxyRows := make([]bool, len(lines))
	galaxyCols := make([]bool, len(lines[0]))
	galaxies = make([]galaxy, 0, 451)
	for row, line := range lines {
		for col, c := range line {
			if c == '#' {
				galaxyCols[col] = true
				galaxyRows[row] = true
			}
		}
	}
	r1 := 0
	c1 := 0
	r2 := 0
	c2 := 0
	for row, line := range lines {
		if !galaxyRows[row] {
			r1 += 2
			r2 += 1_000_000
		} else {
			c1 = 0
			c2 = 0
			for col, c := range line {
				if c == '#' {
					galaxies = append(galaxies, galaxy{r1, c1, r2, c2})
				}
				if !galaxyCols[col] {
					c1 += 2
					c2 += 1_000_000
				} else {
					c1 += 1
					c2 += 1
				}
			}
			r1 += 1
			r2 += 1
		}
	}
	return galaxies
}

func Run(input []byte) (*util.Result, error) {
	start := time.Now()

	galaxies := parseInput(input)

	parse := time.Now()

	part1 := 0
	part2 := 0
	var g1, g2 galaxy
	for gi := 0; gi < len(galaxies)-1; gi++ {
		g1 = galaxies[gi]
		for gj := gi + 1; gj < len(galaxies); gj++ {
			g2 = galaxies[gj]
			dr := g2.row1 - g1.row1
			dc := g2.col1 - g1.col1
			if dr < 0 {
				dr = -dr
			}
			if dc < 0 {
				dc = -dc
			}
			part1 += dr + dc
			dr = g2.row2 - g1.row2
			dc = g2.col2 - g1.col2
			if dr < 0 {
				dr = -dr
			}
			if dc < 0 {
				dc = -dc
			}
			part2 += dr + dc
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
