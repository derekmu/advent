package day06

import (
	"advent/util"
	"bytes"
	_ "embed"
	"math"
	"time"
)

var Problem = util.Problem{Year: "2023", Day: "06", Runner: Run, Input: Input}

//go:embed input.txt
var Input []byte

func parseInput(input []byte) ([]int, int, []int, int) {
	lines := util.ParseInputLines(input)
	times, realTime := parseNums(lines[0])
	dists, realDist := parseNums(lines[1])
	return times, realTime, dists, realDist
}

func Run(input []byte) (*util.Result, error) {
	start := time.Now()

	times, realTime, dists, realDist := parseInput(input)

	parse := time.Now()

	part1 := 1
	for race := 0; race < len(times); race++ {
		t := times[race]
		d := dists[race]
		part1 *= raceWins(t, d)
	}
	part2 := raceWins(realTime, realDist)

	end := time.Now()

	return &util.Result{
		Part1:     part1,
		Part2:     part2,
		StartTime: start,
		ParseTime: parse,
		EndTime:   end,
	}, nil
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
	holdMatch := (float64(-time) + math.Sqrt(float64(time*time-4*dist))) / -2
	// if we use math.Ceil and the match time is exactly an integer, it won't round up and we won't beat the record
	// for example, the math for time=30, dist=200 results in holdMatch=10.0
	// instead, add 1 and floor it to ensure that the hold time is longer than needed to match the record
	hold := int(math.Floor(holdMatch + 1))
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
