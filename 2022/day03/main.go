package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)

	file, err := os.Open("input.txt")
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	sumPriorities1 := 0
	sumPriorities2 := 0
	lines := [3]string{}
	linesI := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// problem 1
		sackOne := line[:len(line)/2]
		sackTwo := line[len(line)/2:]
		sackOneMap := map[int32]bool{}
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
		lines[linesI] = line
		linesI += 1
		if linesI == 3 {
			matchMap := map[int32]int{}
			for _, c := range lines[0] {
				matchMap[c] = 1
			}
			for _, c := range lines[1] {
				if _, ok := matchMap[c]; ok {
					matchMap[c] = 2
				}
			}
			for _, c := range lines[2] {
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
}

func priorityScore(c int32) int {
	if c >= 'A' && c <= 'Z' {
		return int(c - 'A' + 27)
	} else if c >= 'a' && c <= 'z' {
		return int(c - 'a' + 1)
	}
	log.Panicf("Could not determine the score for %c", c)
	return -1
}
