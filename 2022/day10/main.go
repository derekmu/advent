package day10

import (
	"advent/util"
	"bytes"
	"log"
	"time"
)

type op struct {
	name  string
	value int
}

func Run(input []byte) error {
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
	part2 := -1
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
			log.Panicf("unknown operation %s", o.name)
		}
	}

	for _, row := range screen {
		for i := 0; i < len(row); i++ {
			if row[i] {
				print("#")
			} else {
				print(" ")
			}
		}
		println()
	}

	end := time.Now()

	util.PrintResults(part1, part2, start, parse, end)

	return nil
}
