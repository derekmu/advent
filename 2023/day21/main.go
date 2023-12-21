package day21

import (
	"advent/util"
	_ "embed"
	"time"
)

//go:embed input.txt
var Input []byte
var Problem = util.Problem{Year: "2023", Day: "21", Runner: Run, Input: Input}

type point struct {
	row int
	col int
}

var dirs = []point{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func Run(input []byte) (*util.Result, error) {
	start := time.Now()

	lines := util.ParseInputLines(input)

	parse := time.Now()

	part1 := -1
	startPoint := point{}
out:
	for row, line := range lines {
		for col, ch := range line {
			if ch == 'S' {
				startPoint = point{
					row: row,
					col: col,
				}
				break out
			}
		}
	}
	visited := make([][]bool, len(lines))
	for row := range visited {
		visited[row] = make([]bool, len(lines[0]))
	}
	curPoints := make([]point, 0, len(lines)*len(lines[1]))
	curPoints = append(curPoints, startPoint)
	nextPoints := make([]point, 0, len(lines)*len(lines[1]))
	for i := 0; i < 64; i++ {
		for _, p := range curPoints {
			for _, dir := range dirs {
				np := point{row: p.row + dir.row, col: p.col + dir.col}
				if np.row >= 0 && np.row < len(lines) && np.col >= 0 && np.col < len(lines[0]) && lines[np.row][np.col] != '#' && !visited[np.row][np.col] {
					nextPoints = append(nextPoints, np)
					visited[np.row][np.col] = true
				}
			}
		}
		part1 = len(nextPoints)
		for _, np := range nextPoints {
			visited[np.row][np.col] = false
		}
		curPoints, nextPoints = nextPoints, curPoints[:0]
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
