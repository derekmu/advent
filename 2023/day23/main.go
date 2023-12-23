package day23

import (
	"advent/util"
	_ "embed"
	"time"
)

//go:embed input.txt
var Input []byte
var Problem = util.Problem{Year: "2023", Day: "XX", Runner: Run, Input: Input}

type path struct {
	stepCount int
	row       int
	col       int
	visited   [][]bool
}

type direction struct {
	drow  int
	dcol  int
	slope byte
}

var dirs = []direction{
	{dcol: 1, slope: '>'},
	{drow: 1, slope: 'v'},
	{dcol: -1, slope: '<'},
	{drow: -1, slope: '^'},
}

func Run(input []byte) (*util.Result, error) {
	start := time.Now()

	lines := util.ParseInputLines(input)
	width := len(lines)
	height := len(lines[0])

	parse := time.Now()

	empty := make([][]bool, height)
	for i := range empty {
		empty[i] = make([]bool, width)
	}
	part1 := -1
	paths := make([]path, 0, 50)
	paths = append(paths, path{stepCount: 0, row: 0, col: 1, visited: copyVisited(empty)})
	for len(paths) > 0 {
		p := paths[len(paths)-1]
		paths = paths[:len(paths)-1]
		moved := true
		for moved {
			moved = false
			p.visited[p.row][p.col] = true
			if p.row == height-1 && p.col == width-2 {
				part1 = max(part1, p.stepCount)
				break
			}
			switch lines[p.row][p.col] {
			case '>':
				p.col++
				p.stepCount++
				moved = true
				continue
			case '<':
				p.col--
				p.stepCount++
				moved = true
				continue
			case '^':
				p.row--
				p.stepCount++
				moved = true
				continue
			case 'v':
				p.row++
				p.stepCount++
				moved = true
				continue
			}
			op := p
			for _, d := range dirs {
				row := op.row + d.drow
				col := op.col + d.dcol
				if row >= 0 && row < height && col >= 0 && col < width && (lines[row][col] == '.' || lines[row][col] == d.slope) && !p.visited[row][col] {
					if moved {
						paths = append(paths, path{stepCount: op.stepCount + 1, row: row, col: col, visited: copyVisited(p.visited)})
					} else {
						p.row = row
						p.col = col
						p.stepCount++
						moved = true
					}
				}
			}
		}
	}
	part2 := -1
	paths = append(paths, path{stepCount: 0, row: 0, col: 1, visited: copyVisited(empty)})
	for len(paths) > 0 {
		p := paths[len(paths)-1]
		paths = paths[:len(paths)-1]
		moved := true
		for moved {
			moved = false
			p.visited[p.row][p.col] = true
			if p.row == height-1 && p.col == width-2 {
				part2 = max(part2, p.stepCount)
				break
			}
			op := p
			for _, d := range dirs {
				row := op.row + d.drow
				col := op.col + d.dcol
				if row >= 0 && row < height && col >= 0 && col < width && lines[row][col] != '#' && !p.visited[row][col] {
					if moved {
						paths = append(paths, path{stepCount: op.stepCount + 1, row: row, col: col, visited: copyVisited(p.visited)})
					} else {
						p.row = row
						p.col = col
						p.stepCount++
						moved = true
					}
				}
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

func copyVisited(visited [][]bool) [][]bool {
	n := make([][]bool, len(visited))
	for row, visits := range visited {
		n[row] = make([]bool, len(visits))
		copy(n[row], visits)
	}
	return n
}
