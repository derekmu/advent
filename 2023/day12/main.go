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

type index struct {
	si int
	ci int
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
	part2 := 0
	for _, s := range springs {
		cache := make(map[index]int, len(s.status)*len(s.counts))
		part1 += countWays(s.status, 0, s.counts, 0, cache)
		status := make([]byte, len(s.status)*5+4)
		for i, v := range s.status {
			for j := 0; j < 5; j++ {
				status[i+j*len(s.status)+j] = v
			}
		}
		for j := 1; j < 5; j++ {
			status[len(s.status)*j+j-1] = '?'
		}
		counts := make([]int, len(s.counts)*5)
		for i, v := range s.counts {
			for j := 0; j < 5; j++ {
				counts[i+j*len(s.counts)] = v
			}
		}
		cache = make(map[index]int, len(status)*len(counts))
		part2 += countWays(status, 0, counts, 0, cache)
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

func countWays(status []byte, si int, counts []int, ci int, cache map[index]int) int {
	if v, ok := cache[index{si, ci}]; ok {
		return v
	}
	if ci >= len(counts) {
		for si := si; si < len(status); si++ {
			if status[si] == '#' {
				cache[index{si, ci}] = 0
				return 0
			}
		}
		cache[index{si, ci}] = 1
		return 1
	}
	minS := minSpace(counts, ci)
	ways := 0
	for si := si; si <= len(status)-minS; si++ {
		if status[si] == '.' {
			// can't start at a working spring
			continue
		}
		se := si + counts[ci]
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
			ways += countWays(status, se+1, counts, ci+1, cache)
		}
		if status[si] == '#' {
			// can't pass a broken spring
			break
		}
	}
	cache[index{si, ci}] = ways
	return ways
}

func minSpace(counts []int, ci int) int {
	space := 0
	for c2 := ci; c2 < len(counts); c2++ {
		space += counts[c2] + 1
	}
	return space - 1
}

//???##?#???#????? [1 3 2 2 1]
//1.333.22.22.1111 = 4
//1.333.22..22.111 = 3
