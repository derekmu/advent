package day09

import (
	"advent/util"
	_ "embed"
	"time"
)

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
	galaxyCols := make(map[int]bool, 133)
	galaxyRows := make(map[int]bool, 132)
	galaxyCount := 0
	for row, line := range lines {
		for col, c := range line {
			if c == '#' {
				galaxyCount++
				galaxyCols[col] = true
				galaxyRows[row] = true
			}
		}
	}
	galaxies = make([]galaxy, 0, galaxyCount)
	r1 := 0
	c1 := 0
	r2 := 0
	c2 := 0
	for row, line := range lines {
		if _, ok := galaxyRows[row]; !ok {
			r1 += 2
			r2 += 1_000_000
		} else {
			c1 = 0
			c2 = 0
			for col, c := range line {
				if c == '#' {
					galaxies = append(galaxies, galaxy{r1, c1, r2, c2})
				}
				if _, ok := galaxyCols[col]; !ok {
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
	for gi := 0; gi < len(galaxies)-1; gi++ {
		for gj := gi + 1; gj < len(galaxies); gj++ {
			dr := galaxies[gj].row1 - galaxies[gi].row1
			dc := galaxies[gj].col1 - galaxies[gi].col1
			if dr < 0 {
				dr = -dr
			}
			if dc < 0 {
				dc = -dc
			}
			part1 += dr + dc
			dr = galaxies[gj].row2 - galaxies[gi].row2
			dc = galaxies[gj].col2 - galaxies[gi].col2
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
