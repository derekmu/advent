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

	height := len(lines)
	width := len(lines[0])
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
	visited := make([][]bool, height)
	for row := range lines {
		visited[row] = make([]bool, width)
	}
	curPoints := make([]point, 0, height*width)
	curPoints = append(curPoints, startPoint)
	nextPoints := make([]point, 0, height*width)
	for step := 0; step < 64; step++ {
		for _, p := range curPoints {
			for _, dir := range dirs {
				np := point{row: p.row + dir.row, col: p.col + dir.col}
				if np.row >= 0 && np.row < height && np.col >= 0 && np.col < width && lines[np.row][np.col] != '#' && !visited[np.row][np.col] {
					nextPoints = append(nextPoints, np)
					visited[np.row][np.col] = true
				}
			}
		}
		for _, np := range nextPoints {
			visited[np.row][np.col] = false
		}
		curPoints, nextPoints = nextPoints, curPoints[:0]
	}
	part1 := len(curPoints)

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
