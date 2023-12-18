package day10

import (
	"advent/util"
	"bytes"
	_ "embed"
	"time"
)

var Problem = util.Problem{Year: "2023", Day: "10", Runner: Run, Input: Input}

type direction byte

const (
	none direction = iota
	up
	right
	down
	left
)

type point struct {
	row int
	col int
}

type pather struct {
	point
	entered direction
	steps   int
}

func (p pather) goDir(dir direction) pather {
	switch dir {
	case up:
		p.row--
		p.entered = down
	case right:
		p.col++
		p.entered = left
	case down:
		p.row++
		p.entered = up
	case left:
		p.col--
		p.entered = right
	default:
		panic("invalid direction")
	}
	p.steps++
	return p
}

//go:embed input.txt
var Input []byte

func parseInput(input []byte) (lines [][]byte, starter pather) {
	lines = util.ParseInputLines(input)
	for row, line := range lines {
		col := bytes.IndexByte(line, 'S')
		if col >= 0 {
			starter.row = row
			starter.col = col
			break
		}
	}
	return lines, starter
}

func Run(input []byte) (*util.Result, error) {
	start := time.Now()

	lines, starter := parseInput(input)

	parse := time.Now()

	pathers := [2]pather{}
	startDirs := [2]direction{}
	pi := 0
	for dir := up; dir <= left && pi < 2; dir++ {
		p := starter.goDir(dir)
		if p.row >= 0 && p.row < len(lines) && p.col >= 0 && p.col < len(lines[0]) {
			d1, d2 := connections(lines[p.row][p.col])
			if d1 == p.entered || d2 == p.entered {
				pathers[pi] = p
				startDirs[pi] = dir
				pi++
				continue
			}
		}
	}
	loopSet := make(map[point]bool, 13902)
	loopSet[starter.point] = true
	loopSet[pathers[0].point] = true
	loopSet[pathers[1].point] = true
	for pathers[0].row != pathers[1].row || pathers[0].col != pathers[1].col {
		p0d1, p0d2 := connections(lines[pathers[0].row][pathers[0].col])
		if pathers[0].entered != p0d1 {
			pathers[0] = pathers[0].goDir(p0d1)
		} else {
			pathers[0] = pathers[0].goDir(p0d2)
		}
		loopSet[pathers[0].point] = true
		p1d1, p1d2 := connections(lines[pathers[1].row][pathers[1].col])
		if pathers[1].entered != p1d1 {
			pathers[1] = pathers[1].goDir(p1d1)
		} else {
			pathers[1] = pathers[1].goDir(p1d2)
		}
		loopSet[pathers[1].point] = true
	}
	part1 := pathers[0].steps

	part2 := 0
	for row, line := range lines {
		ups := 0
		downs := 0
		inside := false
		for col, c := range line {
			if _, ok := loopSet[point{row, col}]; ok {
				if c == 'S' {
					c = getPipe(startDirs)
				}
				d1, d2 := connections(c)
				if d1 == up {
					ups++
				} else if d1 == down {
					downs++
				}
				if d2 == up {
					ups++
				} else if d2 == down {
					downs++
				}
				for ups > 0 && downs > 0 {
					inside = !inside
					ups--
					downs--
				}
			} else if inside {
				part2++
			}
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

func getPipe(dirs [2]direction) byte {
	if dirs[1] < dirs[0] {
		dirs[0], dirs[1] = dirs[1], dirs[0]
	}
	switch dirs {
	case [2]direction{up, down}:
		return '|'
	case [2]direction{left, right}:
		return '-'
	case [2]direction{up, right}:
		return 'L'
	case [2]direction{up, left}:
		return 'J'
	case [2]direction{down, left}:
		return '7'
	case [2]direction{right, down}:
		return 'F'
	default:
		panic("invalid direction")
	}
}

func connections(c byte) (direction, direction) {
	switch c {
	case '|':
		return up, down
	case '-':
		return left, right
	case 'L':
		return up, right
	case 'J':
		return up, left
	case '7':
		return down, left
	case 'F':
		return right, down
	default:
		return none, none
	}
}
