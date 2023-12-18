package day10

import (
	"advent/util"
	"bytes"
	_ "embed"
	"fmt"
	"time"
)

var Problem = util.Problem{Year: "2022", Day: "10", Runner: Run, Input: Input}

type op struct {
	name  string
	value int
}

//go:embed input.txt
var Input []byte

func Run(input []byte) (*util.Result, error) {
	start := time.Now()

	lines := util.ParseInputLines(input)
	ops := make([]op, 0, len(lines))
	for _, line := range lines {
		name, valueStr, _ := bytes.Cut(line, []byte(" "))
		value := 0
		if bytes.Equal(name, []byte("addx")) {
			value = util.Btoi(valueStr)
		}
		ops = append(ops, op{
			name:  string(name),
			value: value,
		})
	}

	parse := time.Now()

	part1 := 0
	x := 1
	cycle := 0
	screen := [6][40]bool{}
	nextCycle := func() {
		cycle++
		if cycle >= 20 && cycle <= 220 && (cycle-20)%40 == 0 {
			part1 += cycle * x
		}
		row := (cycle - 1) / 40
		col := (cycle - 1) % 40
		screen[row][col] = x >= col-1 && x <= col+1
	}
	for _, o := range ops {
		switch o.name {
		case "noop":
			nextCycle()
		case "addx":
			nextCycle()
			nextCycle()
			x += o.value
		default:
			panic(fmt.Sprintf("unknown operation %s", o.name))
		}
	}

	part2Bytes := make([]byte, len(screen)*(len(screen[0])+1)-1)
	bi := 0
	for ri, row := range screen {
		if ri > 0 {
			part2Bytes[bi] = '\n'
			bi++
		}
		for i := 0; i < len(row); i++ {
			if row[i] {
				part2Bytes[bi] = '#'
			} else {
				part2Bytes[bi] = ' '
			}
			bi++
		}
	}
	part2 := string(part2Bytes)

	end := time.Now()

	return &util.Result{
		Part1:     part1,
		Part2:     part2,
		StartTime: start,
		ParseTime: parse,
		EndTime:   end,
	}, nil
}
