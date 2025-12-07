package day06

import (
	"advent/util"
	_ "embed"
	"log"
	"time"
)

//go:embed input.txt
var Input []byte
var Problem = util.Problem{Year: "2025", Day: "06", Runner: Run, Input: Input}

type op byte

const (
	add op = iota
	mul
)

type num struct {
	val   int
	bytes []byte
}

func parseInput(input []byte) (nums [][]num, ops []op) {
	lines := util.ParseInputLines(input)

	nums = make([][]num, len(lines)-1)
	for i := range nums {
		nums[i] = make([]num, 0, 10)
	}
	opLine := lines[len(lines)-1]
	ops = make([]op, 0, 10)
	for i := 0; i < len(opLine); i++ {
		if opLine[i] == '+' {
			ops = append(ops, add)
		} else if opLine[i] == '*' {
			ops = append(ops, mul)
		} else {
			log.Panicln("invalid operation")
		}
		pi := i
		for i < len(opLine)-1 && opLine[i+1] == ' ' {
			i++
		}
		for j := 0; j < len(nums); j++ {
			var bytes []byte
			if i >= len(opLine)-1 {
				bytes = lines[j][pi:]
			} else {
				bytes = lines[j][pi:i]
			}
			trimmed := bytes
			for trimmed[0] == ' ' {
				trimmed = trimmed[1:]
			}
			for trimmed[len(trimmed)-1] == ' ' {
				trimmed = trimmed[:len(trimmed)-1]
			}
			nums[j] = append(nums[j], num{
				val:   util.Btoi(trimmed),
				bytes: bytes,
			})
		}
	}
	return nums, ops
}

func Run(input []byte) (util.Result, error) {
	start := time.Now()

	nums, ops := parseInput(input)

	parse := time.Now()

	part1 := 0
	part2 := 0

	for i, op := range ops {
		val := nums[0][i].val
		for j := 1; j < len(nums); j++ {
			switch op {
			case add:
				val += nums[j][i].val
			case mul:
				val *= nums[j][i].val
			}
		}
		part1 += val

		val = 0
		for j := 0; j < len(nums[0][i].bytes); j++ {
			v := 0
			m := 1
			for ni := len(nums) - 1; ni >= 0; ni-- {
				ch := nums[ni][i].bytes[j]
				if ch != ' ' {
					v = v + int(ch-'0')*m
					m *= 10
				}
			}
			if val == 0 {
				val = v
			} else {
				switch op {
				case add:
					val += v
				case mul:
					val *= v
				}
			}
		}
		part2 += val
	}

	end := time.Now()

	return util.Result{
		Part1:     part1,
		Part2:     part2,
		StartTime: start,
		ParseTime: parse,
		EndTime:   end,
	}, nil
}
