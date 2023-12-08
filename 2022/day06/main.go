package day06

import (
	"advent/util"
	_ "embed"
	"time"
)

//go:embed input.txt
var Input []byte

func Run(input []byte) (*util.Result, error) {
	start := time.Now()

	lines := util.ParseInputLines(input)

	parse := time.Now()

	part1 := 0
	part2 := 0

	for _, line := range lines {
		chars4 := map[uint8]int{}
		fin4 := false
		chars14 := map[uint8]int{}
		fin14 := false
		for i := 0; i < len(line) && (!fin4 || !fin14); i++ {
			if !fin4 {
				part1, fin4 = updateCharCountMap(chars4, 4, line, i)
			}
			if !fin14 {
				part2, fin14 = updateCharCountMap(chars14, 14, line, i)
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

func updateCharCountMap(charCountMap1 map[uint8]int, size int, line []byte, i int) (int, bool) {
	c := line[i]
	v, _ := charCountMap1[c]
	charCountMap1[c] = v + 1
	if i >= size {
		c1 := line[i-size]
		v1, _ := charCountMap1[c1]
		if v1 == 1 {
			delete(charCountMap1, c1)
		} else {
			charCountMap1[c1] = v1 - 1
		}
	}
	if len(charCountMap1) == size {
		return i + 1, true
	}
	return -1, false
}
