package day21

import (
	"advent/util"
	_ "embed"
	"time"
)

//go:embed input.txt
var Input []byte
var Problem = util.Problem{Year: "2023", Day: "21", Runner: Run, Input: Input}

type coloring byte

const (
	blank coloring = iota
	green
	red
)

// takes 8 minutes
var doPart2 = false

type rayDirection byte

const (
	topRightGoingDown rayDirection = iota
	bottomRightGoingUp
	topLeftGoingDown
	bottomLeftGoingUp
	bottomMiddleGoingUp
	topMiddleGoingDown
	middleRightGoingLeft
	middleLeftGoingRight
)

type ray struct {
	dir   rayDirection
	steps int
	color coloring
}

type colorCount struct {
	reds   int
	greens int
}

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

	part1 := calculatePart1(lines)
	part2 := calculatePart2(lines)

	end := time.Now()

	return &util.Result{
		Part1:     part1,
		Part2:     part2,
		StartTime: start,
		ParseTime: parse,
		EndTime:   end,
	}, nil
}

func calculatePart2(lines [][]byte) int {
	height := len(lines)
	width := len(lines[0])
	// skip sample input with size check
	if width != 131 || height != 131 || !doPart2 {
		return -1
	}
	sp := findStartPoint(lines)
	middleSteps := makeStepCounts(lines, sp, green)
	topLeftSteps := makeStepCounts(lines, point{row: 0, col: 0}, green)
	topRightSteps := makeStepCounts(lines, point{row: 0, col: width - 1}, green)
	bottomLeftSteps := makeStepCounts(lines, point{row: height - 1, col: 0}, green)
	bottomRightSteps := makeStepCounts(lines, point{row: height - 1, col: width - 1}, green)
	bottomMiddleSteps := makeStepCounts(lines, point{row: height - 1, col: sp.col}, red)
	topMiddleSteps := makeStepCounts(lines, point{row: 0, col: sp.col}, red)
	middleRightSteps := makeStepCounts(lines, point{row: sp.row, col: width - 1}, red)
	middleLeftSteps := makeStepCounts(lines, point{row: sp.row, col: 0}, red)
	// initialize with origin cell
	answer := middleSteps[len(middleSteps)-1].reds
	rays := []ray{
		{
			dir:   middleRightGoingLeft,
			steps: sp.col + 1,
			color: green,
		},
		{
			dir:   middleLeftGoingRight,
			steps: width - sp.col,
			color: green,
		},
		{
			dir:   bottomMiddleGoingUp,
			steps: sp.row + 1,
			color: green,
		},
		{
			dir:   topMiddleGoingDown,
			steps: height - sp.row,
			color: green,
		},
	}
	for len(rays) > 0 {
		r := rays[len(rays)-1]
		newColor := r.color
		if newColor == green {
			newColor = red
		} else {
			newColor = green
		}
		rays = rays[:len(rays)-1]
		stepsLeft := 26501365 - r.steps
		if stepsLeft < 0 {
			continue
		}
		var sc []colorCount
		switch r.dir {
		case topRightGoingDown:
			sc = topRightSteps
			rays = append(rays, ray{
				dir:   topRightGoingDown,
				steps: r.steps + height,
				color: newColor,
			})
		case bottomRightGoingUp:
			sc = bottomRightSteps
			rays = append(rays, ray{
				dir:   bottomRightGoingUp,
				steps: r.steps + height,
				color: newColor,
			})
		case topLeftGoingDown:
			sc = topLeftSteps
			rays = append(rays, ray{
				dir:   topLeftGoingDown,
				steps: r.steps + height,
				color: newColor,
			})
		case bottomLeftGoingUp:
			sc = bottomLeftSteps
			rays = append(rays, ray{
				dir:   bottomLeftGoingUp,
				steps: r.steps + height,
				color: newColor,
			})
		case bottomMiddleGoingUp:
			sc = bottomMiddleSteps
			rays = append(rays, ray{
				dir:   bottomMiddleGoingUp,
				steps: r.steps + height,
				color: newColor,
			})
		case topMiddleGoingDown:
			sc = topMiddleSteps
			rays = append(rays, ray{
				dir:   topMiddleGoingDown,
				steps: r.steps + height,
				color: newColor,
			})
		case middleRightGoingLeft:
			sc = middleRightSteps
			rays = append(rays, ray{
				dir:   middleRightGoingLeft,
				steps: r.steps + width,
				color: newColor,
			})
			rays = append(rays, ray{
				dir:   topRightGoingDown,
				steps: r.steps + height - sp.row,
				color: newColor,
			})
			rays = append(rays, ray{
				dir:   bottomRightGoingUp,
				steps: r.steps + sp.row + 1,
				color: newColor,
			})
		case middleLeftGoingRight:
			sc = middleLeftSteps
			rays = append(rays, ray{
				dir:   middleLeftGoingRight,
				steps: r.steps + width,
				color: newColor,
			})
			rays = append(rays, ray{
				dir:   topLeftGoingDown,
				steps: r.steps + height - sp.row,
				color: newColor,
			})
			rays = append(rays, ray{
				dir:   bottomLeftGoingUp,
				steps: r.steps + sp.row + 1,
				color: newColor,
			})
		}
		stepsIn := min(len(sc)-1, stepsLeft)
		switch r.color {
		case green:
			answer += sc[stepsIn].greens
		case red:
			answer += sc[stepsIn].reds
		default:
			panic("unexpected color")
		}
	}
	return answer
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

func makeStepCounts(lines [][]byte, startPoint point, curColor coloring) (stepCounts []colorCount) {
	height := len(lines)
	width := len(lines[0])
	visited := make([][]coloring, height)
	for row := range lines {
		visited[row] = make([]coloring, width)
	}
	visited[startPoint.row][startPoint.col] = curColor
	stepCounts = make([]colorCount, 0, width+height)
	curStepCount := colorCount{}
	if curColor == red {
		curStepCount.reds++
		curColor = green
	} else {
		curStepCount.greens++
		curColor = red
	}
	stepCounts = append(stepCounts, curStepCount)
	curPoints := make([]point, 0, height*width)
	curPoints = append(curPoints, startPoint)
	nextPoints := make([]point, 0, height*width)
	done := false
	for step := 0; !done; step++ {
		done = true
		for _, p := range curPoints {
			for _, dir := range dirs {
				np := point{row: p.row + dir.row, col: p.col + dir.col}
				if np.row >= 0 && np.row < height && np.col >= 0 && np.col < width && lines[np.row][np.col] != '#' && visited[np.row][np.col] == blank {
					visited[np.row][np.col] = curColor
					nextPoints = append(nextPoints, np)
					if curColor == red {
						curStepCount.reds++
					} else {
						curStepCount.greens++
					}
					done = false
				}
			}
		}
		if !done {
			stepCounts = append(stepCounts, curStepCount)
		}
		if curColor == red {
			curColor = green
		} else {
			curColor = red
		}
		curPoints, nextPoints = nextPoints, curPoints[:0]
	}
	//for _, v := range visited {
	//	for _, c := range v {
	//		switch c {
	//		case blank:
	//			print(" ")
	//		case green:
	//			print("G")
	//		case red:
	//			print("R")
	//		}
	//	}
	//	println()
	//}
	//println()
	return stepCounts
}
