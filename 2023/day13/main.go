package template

import (
	"advent/util"
	_ "embed"
	"math/bits"
	"time"
)

var Problem = util.Problem{Year: "2023", Day: "13", Runner: Run, Input: Input}

type rockMap struct {
	rows []uint
	cols []uint
}

//go:embed input.txt
var Input []byte

func parseInput(input []byte) (maps []*rockMap) {
	lines := util.ParseInputLines(input)
	mapLines := make([][]byte, 0, 10)
	maps = make([]*rockMap, 0, 100)
	for li := 0; li <= len(lines); li++ {
		if li == len(lines) || len(lines[li]) == 0 {
			cur := &rockMap{make([]uint, len(mapLines)), make([]uint, len(mapLines[0]))}
			for r := 0; r < len(mapLines); r++ {
				rv := uint(0)
				for c := 0; c < len(mapLines[r]); c++ {
					rv = rv << 1
					if mapLines[r][c] == '#' {
						rv = rv | 1
					}
				}
				cur.rows[r] = rv
			}
			for c := 0; c < len(mapLines[0]); c++ {
				cv := uint(0)
				for r := 0; r < len(mapLines); r++ {
					cv = cv << 1
					if mapLines[r][c] == '#' {
						cv = cv | 1
					}
				}
				cur.cols[c] = cv
			}
			maps = append(maps, cur)
			mapLines = mapLines[:0]
			continue
		} else {
			mapLines = append(mapLines, lines[li])
		}
	}
	return maps
}

func Run(input []byte) (*util.Result, error) {
	start := time.Now()

	rockMaps := parseInput(input)

	parse := time.Now()

	part1 := 0
	part2 := 0
	for _, rm := range rockMaps {
		for r := 0; r < len(rm.rows)-1; r++ {
			ror := uint(0)
			for rl, rr := r, r+1; rl >= 0 && rr < len(rm.rows); rl, rr = rl-1, rr+1 {
				ror |= rm.rows[rl] ^ rm.rows[rr]
			}
			if ror == 0 {
				part1 += (r + 1) * 100
			}
			if bits.OnesCount(ror) == 1 {
				part2 += (r + 1) * 100
			}
		}
		for c := 0; c < len(rm.cols)-1; c++ {
			cor := uint(0)
			for rl, rr := c, c+1; rl >= 0 && rr < len(rm.cols); rl, rr = rl-1, rr+1 {
				cor |= rm.cols[rl] ^ rm.cols[rr]
			}
			if cor == 0 {
				part1 += c + 1
			}
			if bits.OnesCount(cor) == 1 {
				part2 += c + 1
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
