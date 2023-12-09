package day08

import (
	"advent/util"
	_ "embed"
	"fmt"
	"time"
)

type node struct {
	left  uint32
	right uint32
}

var (
	origin      = stoui("AAA")
	destination = stoui("ZZZ")
)

//go:embed input.txt
var Input []byte

func parseInput(input []byte) (moves []byte, nodeMap map[uint32]*node) {
	lines := util.ParseInputLines(input)
	moves = lines[0]
	nodeMap = make(map[uint32]*node, len(lines)-2)
	for _, line := range lines[2:] {
		v := b3toui(line[0:3])
		l := b3toui(line[7:10])
		r := b3toui(line[12:15])
		nodeMap[v] = &node{left: l, right: r}
	}
	return moves, nodeMap
}

func Run(input []byte) (*util.Result, error) {
	start := time.Now()

	moves, nodeMap := parseInput(input)

	parse := time.Now()

	curKey := origin
	part1 := -1
	if cur, ok := nodeMap[curKey]; ok {
		part1 = 0
		for curKey != destination {
			switch moves[part1%len(moves)] {
			case 'L':
				curKey = cur.left
			case 'R':
				curKey = cur.right
			default:
				panic(fmt.Sprintf("unknown move %c", moves[part1%len(moves)]))
			}
			cur = nodeMap[curKey]
			part1++
		}
	}

	steps := make([]int, 0, 6)
	for k, n := range nodeMap {
		if k&0xFF == 'A' {
			step := 0
			for k&0xFF != 'Z' {
				switch moves[step%len(moves)] {
				case 'L':
					k = n.left
				case 'R':
					k = n.right
				default:
					panic(fmt.Sprintf("unknown move %c", moves[step%len(moves)]))
				}
				n = nodeMap[k]
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

func stoui(bytes string) uint32 {
	return uint32(bytes[0])<<16 | uint32(bytes[1])<<8 | uint32(bytes[2])
}

func b3toui(bytes []byte) uint32 {
	return uint32(bytes[0])<<16 | uint32(bytes[1])<<8 | uint32(bytes[2])
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
