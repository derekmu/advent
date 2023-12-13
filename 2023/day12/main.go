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

type cacher struct {
	cache [][]int
	maxCi int
}

func (c *cacher) clear() {
	for i := range c.cache {
		c.cache[i] = c.cache[i][:0]
	}
}

func (c *cacher) value(si, ci int) int {
	if si < len(c.cache) && ci < len(c.cache[si]) {
		return c.cache[si][ci]
	}
	return -1
}

func (c *cacher) set(si, ci int, v int) {
	for si >= len(c.cache) {
		c.cache = append(c.cache, make([]int, 0, c.maxCi))
	}
	for ci >= len(c.cache[si]) {
		c.cache[si] = append(c.cache[si], -1)
	}
	c.cache[si][ci] = v
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
	maxCi := 0
	maxSi := 0
	for _, s := range springs {
		maxSi = max(maxSi, len(s.status)*6+4)
		maxCi = max(maxCi, len(s.counts)*6)
	}

	parse := time.Now()

	cache := &cacher{make([][]int, 0, maxSi), maxCi}
	part1 := 0
	part2 := 0
	for _, s := range springs {
		cache.clear()
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
		cache.clear()
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

func countWays(status []byte, si int, counts []int, ci int, cache *cacher) int {
	if v := cache.value(si, ci); v >= 0 {
		return v
	}
	if ci >= len(counts) {
		for s2 := si; s2 < len(status); s2++ {
			if status[s2] == '#' {
				cache.set(si, ci, 0)
				return 0
			}
		}
		cache.set(si, ci, 1)
		return 1
	}
	minS := minSpace(counts, ci)
	ways := 0
	for s2 := si; s2 <= len(status)-minS; s2++ {
		if status[s2] == '.' {
			// can't start at a working spring
			continue
		}
		se := s2 + counts[ci]
		if se < len(status) && status[se] == '#' {
			if status[s2] == '#' {
				// can't pass a broken spring
				break
			} else {
				// can't end before a broken spring
				continue
			}
		}
		sj := s2 + 1
		for ; sj < se; sj++ {
			if status[sj] == '.' {
				// can't have a working spring in the lot
				break
			}
		}
		if sj == se {
			ways += countWays(status, se+1, counts, ci+1, cache)
		}
		if status[s2] == '#' {
			// can't pass a broken spring
			break
		}
	}
	cache.set(si, ci, ways)
	return ways
}

func minSpace(counts []int, ci int) int {
	space := 0
	for c2 := ci; c2 < len(counts); c2++ {
		space += counts[c2] + 1
	}
	return space - 1
}
