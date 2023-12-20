package day03

import (
	"advent/util"
	_ "embed"
	"fmt"
	"time"
)

//go:embed input.txt
var Input []byte
var Problem = util.Problem{Year: "2022", Day: "03", Runner: Run, Input: Input}

func Run(input []byte) (*util.Result, error) {
	start := time.Now()

	lines := util.ParseInputLines(input)

	parse := time.Now()

	part1 := 0
	part2 := 0
	lines2 := [3][]byte{}
	linesI := 0
	for _, line := range lines {
		// problem 1
		sackOne := line[:len(line)/2]
		sackTwo := line[len(line)/2:]
		sackOneMap := map[byte]bool{}
		for _, c := range sackOne {
			sackOneMap[c] = true
		}
		for _, c := range sackTwo {
			if _, ok := sackOneMap[c]; ok {
				part1 += priorityScore(c)
				break
			}
		}

		// problem 2
		lines2[linesI] = line
		linesI += 1
		if linesI == 3 {
			matchMap := map[byte]int{}
			for _, c := range lines2[0] {
				matchMap[c] = 1
			}
			for _, c := range lines2[1] {
				if _, ok := matchMap[c]; ok {
					matchMap[c] = 2
				}
			}
			for _, c := range lines2[2] {
				if v, ok := matchMap[c]; ok && v == 2 {
					part2 += priorityScore(c)
					break
				}
			}
			linesI = 0
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

func priorityScore(c byte) int {
	if c >= 'A' && c <= 'Z' {
		return int(c - 'A' + 27)
	} else if c >= 'a' && c <= 'z' {
		return int(c - 'a' + 1)
	}
	panic(fmt.Sprintf("Could not determine the score for %c", c))
}
