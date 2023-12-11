package day09

import (
	"advent/util"
	_ "embed"
	"time"
)

type point struct {
	row int
	col int
}

//go:embed input.txt
var Input []byte

func parseInput(input []byte) (space [][]bool, galaxies []point) {
	lines := util.ParseInputLines(input)
	galaxyCols := make(map[int]bool, 50)
	galaxyRows := make(map[int]bool, 50)
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
	space = make([][]bool, 0, len(lines)+len(lines)-len(galaxyRows))
	for row, line := range lines {
		s := make([]bool, 0, len(line)+len(line)-len(galaxyCols))
		for col, c := range line {
			if _, ok := galaxyCols[col]; !ok {
				s = append(s, false)
			}
			if c == '.' {
				s = append(s, false)
			} else {
				s = append(s, true)
			}
		}
		if _, ok := galaxyRows[row]; !ok {
			s2 := make([]bool, len(s))
			space = append(space, s2)
		}
		space = append(space, s)
	}
	galaxies = make([]point, 0, galaxyCount)
	for row, line := range space {
		for col, g := range line {
			if g {
				galaxies = append(galaxies, point{row, col})
			}
		}
	}
	return space, galaxies
}

func Run(input []byte) (*util.Result, error) {
	start := time.Now()

	_, galaxies := parseInput(input)

	parse := time.Now()

	part1 := 0
	for gi := 0; gi < len(galaxies)-1; gi++ {
		for gj := gi + 1; gj < len(galaxies); gj++ {
			dr := galaxies[gj].row - galaxies[gi].row
			dc := galaxies[gj].col - galaxies[gi].col
			if dr < 0 {
				dr = -dr
			}
			if dc < 0 {
				dc = -dc
			}
			part1 += dr + dc
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
