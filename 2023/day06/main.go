package main

import (
	"advent/util"
	"bytes"
	"sort"
	"time"
)

func main() {
	input := util.ReadInput()

	start := time.Now()

	lines := util.ParseInputLines(input)
	times, realTime := parseNums(lines[0])
	dists, realDist := parseNums(lines[1])

	parse := time.Now()

	part1 := 1
	for race := 0; race < len(times); race++ {
		t := times[race]
		d := dists[race]
		part1 *= raceWins(t, d)
	}
	part2 := raceWins(realTime, realDist)

	end := time.Now()

	util.PrintResults(part1, part2, start, parse, end)
}

func raceWins(time, dist int) int {
	hold := sort.Search(time, func(hold int) bool {
		boat := (time - hold) * hold
		return boat > dist
	})
	for ; hold < time; hold++ {
		boat := (time - hold) * hold
		if boat > dist {
			break
		}
	}
	return time - hold - hold + 1
}

func parseNums(line []byte) ([]int, int) {
	nums := make([]int, 0, 4)
	_, line, _ = bytes.Cut(line, []byte(": "))
	totalNum := 0
	var str []byte
	for len(line) > 0 {
		for line[0] == ' ' {
			line = line[1:]
		}
		str, line, _ = bytes.Cut(line, []byte(" "))
		nums = append(nums, util.Btoi(str))
		for _, c := range str {
			totalNum = totalNum*10 + int(c-'0')
		}
	}
	return nums, totalNum
}
