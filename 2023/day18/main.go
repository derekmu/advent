package day18

import (
	"advent/util"
	"bytes"
	_ "embed"
	"image"
	"time"
)

var (
	dirDiffMap = map[byte]image.Point{
		'U': {Y: -1},
		'R': {X: 1},
		'D': {Y: 1},
		'L': {X: -1},
	}
)

var Problem = util.Problem{Year: "2023", Day: "18", Runner: Run, Input: Input}

//go:embed input.txt
var Input []byte

func parseInput(input []byte) (poly1 []image.Point, moves1 int, poly2 []image.Point, moves2 int) {
	lines := util.ParseInputLines(input)
	poly1 = make([]image.Point, 0, len(lines)+1)
	poly2 = make([]image.Point, 0, len(lines)+1)
	var dirBytes, countBytes, colorBytes []byte
	cur1 := image.Point{}
	poly1 = append(poly1, cur1)
	cur2 := image.Point{}
	poly2 = append(poly2, cur2)
	moves1++
	moves2++
	for _, line := range lines {
		dirBytes, line, _ = bytes.Cut(line, []byte(" "))
		countBytes, colorBytes, _ = bytes.Cut(line, []byte(" "))
		count := util.Btoi(countBytes)
		diff := dirDiffMap[dirBytes[0]]
		cur1.Y += count * diff.Y
		cur1.X += count * diff.X
		poly1 = append(poly1, cur1)
		moves1 += count

		count = util.HexBtoi(colorBytes[2:8])
		count = count >> 4
		diff = dirDiffMap[byte(count&0xF)]
		cur2.Y += count * diff.Y
		cur2.X += count * diff.X
		poly2 = append(poly2, cur2)
		moves2 += count
	}
	return poly1, moves1, poly2, moves2
}

func Run(input []byte) (*util.Result, error) {
	start := time.Now()

	poly1, count1, poly2, count2 := parseInput(input)

	parse := time.Now()

	part1 := solve(poly1) + count1/2 + 1
	part2 := solve(poly2) + count2/2 + 1

	end := time.Now()

	return &util.Result{
		Part1:     part1,
		Part2:     part2,
		StartTime: start,
		ParseTime: parse,
		EndTime:   end,
	}, nil
}

func solve(poly []image.Point) int {
	sum := 0
	for i := 0; i < len(poly)-1; i++ {
		m1 := poly[i]
		m2 := poly[i+1]
		sum += (m1.Y + m2.Y) * (m1.X - m2.X)
	}
	if sum < 0 {
		sum = -sum
	}
	return sum / 2
}
