package day14

import (
	"advent/util"
	_ "embed"
	"strings"
	"time"
)

//go:embed input.txt
var Input []byte

func parseInput(input []byte) [][]byte {
	input2 := make([]byte, len(input))
	copy(input2, input)
	return util.ParseInputLines(input2)
}

func Run(input []byte) (*util.Result, error) {
	start := time.Now()

	lines := parseInput(input)

	parse := time.Now()

	part1 := sumLoad1(lines)
	patterns := make(map[string]int, 1_000)
	patterns[makePattern(lines)] = 0
	var c1 int
	c := 0
	for ; c < 1_000_000_000; c++ {
		spinCycle(lines)
		pattern := makePattern(lines)
		if c2, ok := patterns[pattern]; ok {
			c1 = c2
			break
		} else {
			patterns[pattern] = c
		}
	}
	cd := c - c1
	c = (1_000_000_000-c1)/cd*cd + c1 + 1
	for ; c < 1_000_000_000; c++ {
		spinCycle(lines)
	}
	part2 := sumLoad2(lines)

	end := time.Now()

	return &util.Result{
		Part1:     part1,
		Part2:     part2,
		StartTime: start,
		ParseTime: parse,
		EndTime:   end,
	}, nil
}

func makePattern(lines [][]byte) string {
	var builder strings.Builder
	builder.Grow(len(lines) * len(lines[0]))
	for _, line := range lines {
		builder.Write(line)
	}
	return builder.String()
}

func spinCycle(lines [][]byte) {
	rows := len(lines)
	cols := len(lines[0])
	// north
	for c := 0; c < cols; c++ {
		r1 := 0
		for r := 0; r < rows; r++ {
			switch lines[r][c] {
			case 'O':
				lines[r][c] = '.'
				lines[r1][c] = 'O'
				r1++
			case '#':
				r1 = r + 1
			case '.':
			default:
				panic("invalid character")
			}
		}
	}
	// west
	for r := 0; r < rows; r++ {
		c1 := 0
		for c := 0; c < cols; c++ {
			switch lines[r][c] {
			case 'O':
				lines[r][c] = '.'
				lines[r][c1] = 'O'
				c1++
			case '#':
				c1 = c + 1
			case '.':
			default:
				panic("invalid character")
			}
		}
	}
	// south
	for c := 0; c < cols; c++ {
		r1 := rows - 1
		for r := rows - 1; r >= 0; r-- {
			switch lines[r][c] {
			case 'O':
				lines[r][c] = '.'
				lines[r1][c] = 'O'
				r1--
			case '#':
				r1 = r - 1
			case '.':
			default:
				panic("invalid character")
			}
		}
	}
	// east
	for r := 0; r < rows; r++ {
		c1 := cols - 1
		for c := cols - 1; c >= 0; c-- {
			switch lines[r][c] {
			case 'O':
				lines[r][c] = '.'
				lines[r][c1] = 'O'
				c1--
			case '#':
				c1 = c - 1
			case '.':
			default:
				panic("invalid character")
			}
		}
	}
}

func sumLoad1(lines [][]byte) int {
	load := 0
	cols := len(lines[0])
	for c := 0; c < cols; c++ {
		r1 := 0
		for r, line := range lines {
			switch line[c] {
			case 'O':
				load += cols - r1
				r1++
			case '#':
				r1 = r + 1
			case '.':
			default:
				panic("invalid character")
			}
		}
	}
	return load
}

func sumLoad2(lines [][]byte) int {
	load := 0
	cols := len(lines[0])
	for c := 0; c < cols; c++ {
		for r, line := range lines {
			switch line[c] {
			case 'O':
				load += cols - r
			case '#':
			case '.':
			default:
				panic("invalid character")
			}
		}
	}
	return load
}
