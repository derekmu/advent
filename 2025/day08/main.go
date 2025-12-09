package day08

import (
	"advent/util"
	_ "embed"
	"slices"
	"time"
)

//go:embed input.txt
var Input []byte
var Problem = util.Problem{Year: "2025", Day: "08", Runner: Run, Input: Input}

type box struct {
	x, y, z int
	circuit *circuit
}

type distance struct {
	bi1, bi2 int
	distance float64
}

type circuit struct {
	bis     []int
	visited bool
}

func parseInput(input []byte) (connectCount int, boxes []box) {
	lines := util.ParseInputLines(input)
	connectCount = util.Btoi(lines[0])
	lines = lines[1:]
	boxes = make([]box, len(lines))
	for i, line := range lines {
		parts := util.ParseInputDelimiter(line, []byte(","))
		boxes[i] = box{
			x:       util.Btoi(parts[0]),
			y:       util.Btoi(parts[1]),
			z:       util.Btoi(parts[2]),
			circuit: &circuit{bis: []int{i}},
		}
	}
	return connectCount, boxes
}

func Run(input []byte) (util.Result, error) {
	start := time.Now()

	connectCount, boxes := parseInput(input)

	parse := time.Now()

	part1 := -1
	part2 := -1

	distances := make([]distance, 0, len(boxes)*(len(boxes)-1)/2)
	for bi1 := range boxes {
		for bi2 := bi1 + 1; bi2 < len(boxes); bi2++ {
			dx := float64(boxes[bi1].x - boxes[bi2].x)
			dy := float64(boxes[bi1].y - boxes[bi2].y)
			dz := float64(boxes[bi1].z - boxes[bi2].z)
			distances = append(distances, distance{
				bi1:      bi1,
				bi2:      bi2,
				distance: dx*dx + dy*dy + dz*dz,
			})
		}
	}
	slices.SortFunc(distances, func(a, b distance) int {
		if a.distance < b.distance {
			return -1
		} else if b.distance < a.distance {
			return 1
		} else {
			return 0
		}
	})
	di := 0
	for ; di < connectCount; di++ {
		connect(boxes, distances[di].bi1, distances[di].bi2)
	}
	s1, s2, s3 := 0, 0, 0
	for bi := range boxes {
		c := boxes[bi].circuit
		if !c.visited {
			c.visited = true
			s := len(c.bis)
			if s > s1 {
				s3 = s2
				s2 = s1
				s1 = s
			} else if s > s2 {
				s3 = s2
				s2 = s
			} else if s > s3 {
				s3 = s
			}
		}
	}
	part1 = s1 * s2 * s3
	for ; di < len(distances) && part2 < 0; di++ {
		part2 = connect(boxes, distances[di].bi1, distances[di].bi2)
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

func connect(boxes []box, bi1 int, bi2 int) (part2 int) {
	c1 := boxes[bi1].circuit
	c2 := boxes[bi2].circuit
	if c1 == c2 {
		return -1
	}
	// combine all the boxes in both circuits
	bis := make([]int, 0, len(c1.bis)+len(c2.bis))
	bis = append(bis, c1.bis...)
	bis = append(bis, c2.bis...)
	if len(bis) == len(boxes) {
		return boxes[bi1].x * boxes[bi2].x
	}
	// update box 1's circuit
	c1.bis = bis
	// for everything in box 2's circuit, move them to box 1's circuit
	for _, bi := range c2.bis {
		boxes[bi].circuit = c1
	}
	return -1
}
