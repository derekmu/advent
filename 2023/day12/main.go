package day12

import (
	"advent/util"
	"bytes"
	_ "embed"
	"time"
)

type spring struct {
	status []byte
	counts []int
}

//go:embed input.txt
var Input []byte

func parseInput(input []byte) (springs []*spring) {
	lines := util.ParseInputLines(input)
	springs = make([]*spring, 0, len(lines))
	for _, line := range lines {
		status, nums, _ := bytes.Cut(line, []byte(" "))
		countsBytes := bytes.Split(nums, []byte(","))
		counts := make([]int, len(countsBytes))
		for i, b := range countsBytes {
			counts[i] = util.Btoi(b)
		}
		springs = append(springs, &spring{
			status: status,
			counts: counts,
		})
	}
	return springs
}

func Run(input []byte) (*util.Result, error) {
	start := time.Now()

	springs := parseInput(input)

	parse := time.Now()

	part1 := 0
	for _, s := range springs {
		part1 += countWays(s.status, s.counts)
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

func countWays(status []byte, counts []int) int {
	if len(counts) == 0 {
		for _, s := range status {
			if s == '#' {
				// haven't consumed all broken springs
				return 0
			}
		}
		return 1
	}
	minS := minSpace(counts)
	ways := 0
	for si := 0; si <= len(status)-minS; si++ {
		if status[si] == '.' {
			// can't start at a working spring
			continue
		}
		se := si + counts[0]
		if se < len(status) && status[se] == '#' {
			if status[si] == '#' {
				// can't pass a broken spring
				break
			} else {
				// can't end before a broken spring
				continue
			}
		}
		sj := si + 1
		for ; sj < se; sj++ {
			if status[sj] == '.' {
				// can't have a working spring in the lot
				break
			}
		}
		if sj == se {
			ways += countWays(status[min(se+1, len(status)):], counts[1:])
		}
		if status[si] == '#' {
			// can't pass a broken spring
			break
		}
	}
	return ways
}

func minSpace(counts []int) int {
	space := 0
	for _, c := range counts {
		space += c + 1
	}
	return space - 1
}

//???##?#???#????? [1 3 2 2 1]
//1.333.22.22.1111 = 4
//1.333.22..22.111 = 3
