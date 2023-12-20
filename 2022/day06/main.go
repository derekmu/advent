package day06

import (
	"advent/util"
	_ "embed"
	"time"
)

//go:embed input.txt
var Input []byte
var Problem = util.Problem{Year: "2022", Day: "06", Runner: Run, Input: Input}

func Run(input []byte) (*util.Result, error) {
	start := time.Now()

	lines := util.ParseInputLines(input)

	parse := time.Now()

	part1 := make([]int, 0, len(lines))
	part2 := make([]int, 0, len(lines))

	var p1, p2 int
	for _, line := range lines {
		chars4 := map[uint8]int{}
		fin4 := false
		chars14 := map[uint8]int{}
		fin14 := false
		for i := 0; i < len(line) && (!fin4 || !fin14); i++ {
			if !fin4 {
				p1, fin4 = updateCharCountMap(chars4, 4, line, i)
			}
			if !fin14 {
				p2, fin14 = updateCharCountMap(chars14, 14, line, i)
			}
		}
		part1 = append(part1, p1)
		part2 = append(part2, p2)
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
