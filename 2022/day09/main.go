package day09

import (
	"advent/util"
	"log"
	"time"
)

type move struct {
	dir  byte
	dist int
}

type point struct {
	x int
	y int
}

func (p *point) update(f *point) bool {
	if abs(f.x-p.x) > 1 || abs(f.y-p.y) > 1 {
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

func Run(input []byte) error {
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
	end := time.Now()

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

	util.PrintResults(part1, part2, start, parse, end)

	return nil
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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
		log.Panicf("unknown direction %c", dir)
		return 0, 0
	}
}
