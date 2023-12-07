package day06

import (
	"advent/util"
	"bytes"
	"math"
	"time"
)

func Run(input []byte) error {
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

	return nil
}

func raceWins(time, dist int) int {
	// formula for the time taken to finish:
	//	(t-h)*h
	//	-h^2 + th
	// where we match the record:
	//	d = -h^2 + th
	// 	0 = -h^2 + th - d
	// apply the quadratic equation:
	// 	h = (-b +- sqrt(b^2 - 4ac)) / 2a
	//  a = -1
	//  b = t
	//  c = -d
	hold := int(math.Ceil((float64(-time) + math.Sqrt(float64(time*time-4*dist))) / -2))
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
