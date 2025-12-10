package day09

import (
	"advent/util"
	_ "embed"
	"slices"
	"time"
)

//go:embed input.txt
var Input []byte
var Problem = util.Problem{Year: "2025", Day: "09", Runner: Run, Input: Input}

type tile byte

const (
	outside tile = iota
	edge
	inside
)

type point struct {
	x, y int
}

func parseInput(input []byte) (points []point) {
	lines := util.ParseInputLines(input)
	points = make([]point, len(lines))
	for i, line := range lines {
		parts := util.ParseInputDelimiter(line, []byte(","))
		points[i] = point{
			x: util.Btoi(parts[0]),
			y: util.Btoi(parts[1]),
		}
	}
	return points
}

func Run(input []byte) (util.Result, error) {
	start := time.Now()

	points := parseInput(input)

	parse := time.Now()

	part1 := -1
	part2 := -1

	xs := make([]int, 0, len(points))
	ys := make([]int, 0, len(points))
	xm := make(map[int]int, len(points))
	ym := make(map[int]int, len(points))
	for i := range points {
		if _, exists := xm[points[i].x]; !exists {
			xs = append(xs, points[i].x)
			xm[points[i].x] = -1
		}
		if _, exists := ym[points[i].y]; !exists {
			ys = append(ys, points[i].y)
			ym[points[i].y] = -1
		}
	}
	slices.Sort(xs)
	slices.Sort(ys)
	for i, x := range xs {
		xm[x] = i
	}
	for i, y := range ys {
		ym[y] = i
	}
	tiles := make([][]tile, len(xs))
	for x := range tiles {
		tiles[x] = make([]tile, len(ys))
	}
	for i := 0; i < len(points); i++ {
		p1 := points[i]
		p2 := points[(i+1)%len(points)]
		x1 := xm[min(p1.x, p2.x)]
		x2 := xm[max(p1.x, p2.x)]
		y1 := ym[min(p1.y, p2.y)]
		y2 := ym[max(p1.y, p2.y)]
		for x := x1; x <= x2; x++ {
			for y := y1; y <= y2; y++ {
				tiles[x][y] = edge
			}
		}
	}
	for x, col := range tiles {
		in := false
		for y := 0; y < len(col); y++ {
			if col[y] == edge {
				// if we have a sequence of edges, we're travelling down a vertical edge
				if y < len(col)-1 && col[y+1] == edge {
					inRight := x < len(tiles)-1 && tiles[x+1][y] == edge
					for y < len(col) && col[y] == edge {
						y++
					}
					y--
					outRight := x < len(tiles)-1 && tiles[x+1][y] == edge
					// if the horizontal edges are on opposite sides of the column, we've switched from inside to outside or vice versa
					// otherwise, only the edge is inside, we don't switch
					if inRight != outRight {
						in = !in
					}
				} else {
					// we hit a 1 tile edge, always switch
					in = !in
				}
			} else if in {
				col[y] = inside
			}
		}
	}
	//for y := range tiles[0] {
	//	for x := range tiles {
	//		switch tiles[x][y] {
	//		case outside:
	//			print(" ")
	//		case edge:
	//			print("#")
	//		case inside:
	//			print(".")
	//		}
	//	}
	//	println()
	//}
	for i := 0; i < len(points); i++ {
		p1 := points[i]
		for j := i + 1; j < len(points); j++ {
			p2 := points[j]
			dx := util.Abs(p1.x-p2.x) + 1
			dy := util.Abs(p1.y-p2.y) + 1
			if dx*dy > part1 {
				part1 = dx * dy
			}
			x1 := xm[min(p1.x, p2.x)]
			x2 := xm[max(p1.x, p2.x)]
			y1 := ym[min(p1.y, p2.y)]
			y2 := ym[max(p1.y, p2.y)]
			good := true
		out:
			for x := x1; x <= x2; x++ {
				for y := y1; y <= y2; y++ {
					if tiles[x][y] == outside {
						good = false
						break out
					}
				}
			}
			if good && dx*dy > part2 {
				part2 = dx * dy
			}
		}
	}

	end := time.Now()

	return util.Result{
		Part1:     part1,
		Part2:     part2,
		StartTime: start,
		ParseTime: parse,
		EndTime:   end,
	}, nil
}
