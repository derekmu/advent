package day09

import (
	"advent/util"
	_ "embed"
	"fmt"
	"time"
)

//go:embed input.txt
var Input []byte
var Problem = util.Problem{Year: "2022", Day: "09", Runner: Run, Input: Input}

type move struct {
	dir  byte
	dist int
}

type point struct {
	x int
	y int
}

func (p *point) update(f *point) bool {
	if util.Abs(f.x-p.x) > 1 || util.Abs(f.y-p.y) > 1 {
		if f.x > p.x {
			p.x += 1
		} else if f.x < p.x {
			p.x -= 1
		}
		if f.y > p.y {
			p.y += 1
		} else if f.y < p.y {
			p.y -= 1
		}
		return true
	}
	return false
}

func Run(input []byte) (util.Result, error) {
	start := time.Now()

	lines := util.ParseInputLines(input)
	moves := make([]move, 0, len(lines))
	for _, line := range lines {
		moves = append(moves, move{
			dir:  line[0],
			dist: util.Btoi(line[2:]),
		})
	}

	parse := time.Now()

	part1 := 0
	h := point{}
	t := point{}
	visited := make(map[point]bool, 6209)
	visited[t] = true
	for _, m := range moves {
		dx, dy := dxy(m.dir)
		for i := 0; i < m.dist; i++ {
			h.x += dx
			h.y += dy
			if t.update(&h) {
				visited[t] = true
			}
		}
	}
	part1 = len(visited)

	part2 := 0
	knots := make([]point, 10)
	visited = make(map[point]bool, 10000)
	visited[knots[9]] = true
	for _, m := range moves {
		dx, dy := dxy(m.dir)
		for i := 0; i < m.dist; i++ {
			knots[0].x += dx
			knots[0].y += dy
			for ki := 1; ki < len(knots); ki++ {
				if !knots[ki].update(&knots[ki-1]) {
					break
				} else if ki == 9 {
					visited[knots[ki]] = true
				}
			}
		}
	}
	part2 = len(visited)

	end := time.Now()

	return util.Result{
		Part1:     part1,
		Part2:     part2,
		StartTime: start,
		ParseTime: parse,
		EndTime:   end,
	}, nil
}

func dxy(dir byte) (int, int) {
	switch dir {
	case 'U':
		return 0, 1
	case 'R':
		return 1, 0
	case 'D':
		return 0, -1
	case 'L':
		return -1, 0
	default:
		panic(fmt.Sprintf("unknown direction %c", dir))
	}
}
