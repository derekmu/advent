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

func Run(input []byte) (util.Result, error) {
	start := time.Now()

	lines := util.ParseInputLines(input)

	parse := time.Now()

	part1 := calculatePart1(lines)
	part2 := calculatePart2(lines)

	end := time.Now()

	return util.Result{
		Part1:     part1,
		Part2:     part2,
		StartTime: start,
		ParseTime: parse,
		EndTime:   end,
	}, nil
}

func findStartPoint(lines [][]byte) point {
	for row, line := range lines {
		for col, ch := range line {
			if ch == 'S' {
				return point{
					row: row,
					col: col,
				}
			}
		}
	}
	panic("start point not found")
}

func calculatePart1(lines [][]byte) int {
	height := len(lines)
	width := len(lines[0])
	visited := make([][]bool, height)
	for row := range lines {
		visited[row] = make([]bool, width)
	}
	curPoints := make([]point, 0, height*width)
	curPoints = append(curPoints, findStartPoint(lines))
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
	return len(curPoints)
}

func simulate(lines [][]byte, steps []int) []int {
	height := len(lines)
	width := len(lines[0])
	result := make([]int, 0, len(steps))
	stepI := 0
	curPoints := make([]point, 0, height*width)
	curPoints = append(curPoints, findStartPoint(lines))
	nextPoints := make([]point, 0, height*width)
	for step := 0; step < steps[len(steps)-1]; step++ {
		visited := make(map[point]bool, height)
		for _, p := range curPoints {
			for _, dir := range dirs {
				np := point{row: p.row + dir.row, col: p.col + dir.col}
				if lines[mod(np.row, height)][mod(np.col, width)] != '#' {
					if _, ok := visited[np]; !ok {
						nextPoints = append(nextPoints, np)
						visited[np] = true
					}
				}
			}
		}
		curPoints, nextPoints = nextPoints, curPoints[:0]
		if step == steps[stepI]-1 {
			result = append(result, len(curPoints))
			stepI++
		}
	}
	return result
}

func calculatePart2(lines [][]byte) int {
	// don't calculate for sample inputs
	if len(lines) != 131 {
		return -1
	}
	x := 26_501_365 / len(lines)
	result := simulate(lines, []int{65, 65 + 131, 65 + 131 + 131})
	a := result[0]
	b := result[1]
	c := result[2]
	return a + x*(b-a) + x*(x-1)/2*((c-b)-(b-a))
}

func mod(a, b int) int {
	return (a%b + b) % b
}
