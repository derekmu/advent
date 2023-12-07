package day03

import (
	"advent/util"
	"time"
)

type number struct {
	row      int
	start    int
	end      int
	adjacent bool
}

type index struct {
	row int
	col int
}

func Run(input []byte) error {
	start := time.Now()

	lines := util.ParseInputLines(input)
	numbers := make([]*number, 0, 1207)
	gearMap := make(map[index][]int, 368)
	for row, line := range lines {
		lines = append(lines, line)
		var currentNumber *number
		for col, ch := range line {
			if ch >= '0' && ch <= '9' {
				if currentNumber != nil && currentNumber.end == col-1 {
					currentNumber.end = col
				} else {
					currentNumber = &number{
						row:   row,
						start: col,
						end:   col,
					}
					numbers = append(numbers, currentNumber)
				}
			} else if ch == '*' {
				gearMap[index{row: row, col: col}] = make([]int, 0, 2)
			}
		}
	}

	parse := time.Now()

	part1 := 0
	for _, n := range numbers {
		for row := n.row - 1; row <= n.row+1; row++ {
			if row < 0 || row >= len(lines) {
				continue
			}
			for col := n.start - 1; col <= n.end+1; col++ {
				if row == n.row && col >= n.start && col <= n.end {
					col = n.end + 1
				}
				if col < 0 || col >= len(lines[row]) {
					continue
				}
				ch := lines[row][col]
				switch ch {
				case '.', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				default:
					i := util.Btoi(lines[n.row][n.start : n.end+1])
					part1 += i
					if ch == '*' {
						g := gearMap[index{row: row, col: col}]
						gearMap[index{row: row, col: col}] = append(g, i)
					}
				}
			}
		}
	}

	part2 := 0
	for _, g := range gearMap {
		if len(g) == 2 {
			part2 += g[0] * g[1]
		}
	}

	end := time.Now()

	util.PrintResults(part1, part2, start, parse, end)

	return nil
}
