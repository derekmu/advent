package day18

import (
	"advent/util"
	"bytes"
	_ "embed"
	"image"
	"time"
)

type direction byte

const (
	up direction = iota
	right
	down
	left
)

type tile byte

const (
	empty tile = iota
	trench
	outside
)

var (
	charDirMap = map[byte]direction{
		'U': up,
		'R': right,
		'D': down,
		'L': left,
	}
	dirDiffMap = map[direction]image.Point{
		up:    {Y: -1},
		right: {X: 1},
		down:  {Y: 1},
		left:  {X: -1},
	}
)

var Problem = util.Problem{Year: "2023", Day: "18", Runner: Run, Input: Input}

type move struct {
	dir   direction
	count int
	color []byte
}

//go:embed input.txt
var Input []byte

func parseInput(input []byte) (moves []move, size image.Rectangle) {
	lines := util.ParseInputLines(input)
	moves = make([]move, 0, len(lines))
	var dirBytes, countBytes, colorBytes []byte
	cur := image.Point{}
	for _, line := range lines {
		dirBytes, line, _ = bytes.Cut(line, []byte(" "))
		countBytes, colorBytes, _ = bytes.Cut(line, []byte(" "))
		m := move{
			dir:   charDirMap[dirBytes[0]],
			count: util.Btoi(countBytes),
			color: colorBytes[2:8],
		}
		moves = append(moves, m)
		diff := dirDiffMap[m.dir]
		cur.Y += m.count * diff.Y
		cur.X += m.count * diff.X
		size.Min.X = min(size.Min.X, cur.X)
		size.Min.Y = min(size.Min.Y, cur.Y)
		size.Max.X = max(size.Max.X, cur.X)
		size.Max.Y = max(size.Max.Y, cur.Y)
	}
	return moves, size
}

func Run(input []byte) (*util.Result, error) {
	start := time.Now()

	moves, size := parseInput(input)

	parse := time.Now()

	tiles := make([][]tile, size.Dy()+3)
	for row := range tiles {
		tiles[row] = make([]tile, size.Dx()+3)
	}
	cur := image.Point{X: max(0, -size.Min.X) + 1, Y: max(0, -size.Min.Y) + 1}
	tiles[cur.Y][cur.X] = trench
	for _, m := range moves {
		diff := dirDiffMap[m.dir]
		for c := 1; c <= m.count; c++ {
			tiles[cur.Y+c*diff.Y][cur.X+c*diff.X] = trench
		}
		cur.Y += m.count * diff.Y
		cur.X += m.count * diff.X
	}
	fillStack := make([]image.Point, 0, 100)
	addOutside := func(p image.Point) {
		if p.Y >= 0 && p.X >= 0 && p.Y < len(tiles) && p.X < len(tiles[0]) && tiles[p.Y][p.X] == empty {
			tiles[p.Y][p.X] = outside
			fillStack = append(fillStack, p)
		}
	}
	addOutside(image.Point{})
	for len(fillStack) > 0 {
		p := fillStack[len(fillStack)-1]
		fillStack = fillStack[:len(fillStack)-1]
		addOutside(image.Point{X: p.X + 1, Y: p.Y})
		addOutside(image.Point{X: p.X - 1, Y: p.Y})
		addOutside(image.Point{X: p.X, Y: p.Y + 1})
		addOutside(image.Point{X: p.X, Y: p.Y - 1})
	}
	part1 := 0
	for _, rowTiles := range tiles {
		for _, t := range rowTiles {
			if t != outside {
				part1++
			}
		}
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
