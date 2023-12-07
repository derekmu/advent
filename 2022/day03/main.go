package day03

import (
	"advent/util"
	"log"
)

func Run(input []byte) error {
	sumPriorities1 := 0
	sumPriorities2 := 0
	lines2 := [3][]byte{}
	linesI := 0

	lines := util.ParseInputLines(input)
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
				sumPriorities1 += priorityScore(c)
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
					sumPriorities2 += priorityScore(c)
					break
				}
			}
			linesI = 0
		}
	}

	log.Printf("The sum of the priorities of misstored items is %d.", sumPriorities1)
	log.Printf("The sum of the priorities of badges is %d.", sumPriorities2)

	return nil
}

func priorityScore(c byte) int {
	if c >= 'A' && c <= 'Z' {
		return int(c - 'A' + 27)
	} else if c >= 'a' && c <= 'z' {
		return int(c - 'a' + 1)
	}
	log.Panicf("Could not determine the score for %c", c)
	return -1
}
