package day22

import (
	"advent/util"
	"bytes"
	_ "embed"
	"slices"
	"time"
)

//go:embed input.txt
var Input []byte
var Problem = util.Problem{Year: "2023", Day: "22", Runner: Run, Input: Input}

type point struct {
	x int
	y int
	z int
}

type brick struct {
	min        point
	max        point
	supporting []*brick
	supporters []*brick
}

func parseInput(input []byte) (bricks []*brick) {
	lines := util.ParseInputLines(input)
	bricks = make([]*brick, len(lines))
	var x1, y1, z1, x2, y2, z2 []byte
	for i, line := range lines {
		x1, line, _ = bytes.Cut(line, []byte(","))
		y1, line, _ = bytes.Cut(line, []byte(","))
		z1, line, _ = bytes.Cut(line, []byte("~"))
		x2, line, _ = bytes.Cut(line, []byte(","))
		y2, z2, _ = bytes.Cut(line, []byte(","))
		x1i := util.Btoi(x1)
		y1i := util.Btoi(y1)
		z1i := util.Btoi(z1)
		x2i := util.Btoi(x2)
		y2i := util.Btoi(y2)
		z2i := util.Btoi(z2)
		bricks[i] = &brick{
			min: point{
				x: min(x1i, x2i),
				y: min(y1i, y2i),
				z: min(z1i, z2i),
			},
			max: point{
				x: max(x1i, x2i),
				y: max(y1i, y2i),
				z: max(z1i, z2i),
			},
		}
	}
	return bricks
}

func Run(input []byte) (*util.Result, error) {
	start := time.Now()

	bricks := parseInput(input)

	parse := time.Now()

	slices.SortFunc(bricks, func(a, b *brick) int {
		if a.min.z < b.min.z {
			return -1
		} else if a.min.z > b.min.z {
			return 1
		} else {
			return 0
		}
	})
	for i, b := range bricks {
		nz := 1
		b.supporters = make([]*brick, 0, 4)
		b.supporting = make([]*brick, 0, 4)
		for j := i - 1; j >= 0; j-- {
			b2 := bricks[j]
			if b.min.x <= b2.max.x && b.max.x >= b2.min.x && b.min.y <= b2.max.y && b.max.y >= b2.min.y {
				nz2 := b2.max.z + 1
				if nz2 == nz {
					b.supporters = append(b.supporters, b2)
				} else if nz2 > nz {
					nz = nz2
					b.supporters = b.supporters[:0]
					b.supporters = append(b.supporters, b2)
				}
			}
		}
		dz := b.min.z - nz
		b.min.z = nz
		b.max.z -= dz
		for _, sb := range b.supporters {
			sb.supporting = append(sb.supporting, b)
		}
	}
	part1 := 0
	for _, b := range bricks {
		good := true
		for _, sb := range b.supporting {
			if len(sb.supporters) <= 1 {
				good = false
				break
			}
		}
		if good {
			part1++
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
