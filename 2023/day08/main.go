package day07

import (
	"advent/util"
	_ "embed"
	"fmt"
	"time"
)

type node struct {
	value uint32
	left  uint32
	right uint32
}

const (
	origin      = uint32('A')<<16 | uint32('A')<<8 | uint32('A')
	destination = uint32('Z')<<16 | uint32('Z')<<8 | uint32('Z')
)

//go:embed input.txt
var Input []byte

func Run(input []byte) (*util.Result, error) {
	start := time.Now()

	lines := util.ParseInputLines(input)
	moves := lines[0]
	nodeMap := make(map[uint32]*node, len(lines)-2)
	for _, line := range lines[2:] {
		v := uint32(line[0])<<16 | uint32(line[1])<<8 | uint32(line[2])
		l := uint32(line[7])<<16 | uint32(line[8])<<8 | uint32(line[9])
		r := uint32(line[12])<<16 | uint32(line[13])<<8 | uint32(line[14])
		nodeMap[v] = &node{value: v, left: l, right: r}
	}

	parse := time.Now()

	cur := nodeMap[origin]
	part1 := 0
	for cur.value != destination {
		switch moves[part1%len(moves)] {
		case 'L':
			cur = nodeMap[cur.left]
		case 'R':
			cur = nodeMap[cur.right]
		default:
			panic(fmt.Sprintf("unknown move %c", moves[part1%len(moves)]))
		}
		part1++
	}

	steps := make([]int, 0, 6)
	for _, n := range nodeMap {
		if n.value&0xFF == 'A' {
			step := 0
			for n.value&0xFF != 'Z' {
				switch moves[step%len(moves)] {
				case 'L':
					n = nodeMap[n.left]
				case 'R':
					n = nodeMap[n.right]
				default:
					panic(fmt.Sprintf("unknown move %c", moves[step%len(moves)]))
				}
				step++
			}
			steps = append(steps, step)
		}
	}
	part2 := findLCM(steps)

	end := time.Now()

	return &util.Result{
		Part1:     part1,
		Part2:     part2,
		StartTime: start,
		ParseTime: parse,
		EndTime:   end,
	}, nil
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func findLCM(numbers []int) int {
	result := numbers[0]
	for i := 1; i < len(numbers); i++ {
		result = lcm(result, numbers[i])
	}
	return result
}
