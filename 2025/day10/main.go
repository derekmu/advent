package day10

import (
	"advent/util"
	"bytes"
	_ "embed"
	"math/bits"
	"time"
)

//go:embed input.txt
var Input []byte
var Problem = util.Problem{Year: "2025", Day: "10", Runner: Run, Input: Input}

type machine struct {
	lights   uint64
	buttons  []uint64
	joltages []int
}

func parseInput(input []byte) (machines []machine) {
	lines := util.ParseInputLines(input)
	machines = make([]machine, len(lines))
	for i, line := range lines {
		lightsLine := line[1:bytes.IndexByte(line, ']')]
		lights := uint64(0)
		for i, b := range lightsLine {
			if b == '#' {
				lights |= 1 << i
			}
		}

		buttonsLine := line[bytes.IndexByte(line, '(') : bytes.LastIndexByte(line, ')')+1]
		parts := util.ParseInputDelimiter(buttonsLine, []byte(" "))
		buttons := make([]uint64, len(parts))
		for i, p := range parts {
			buttonParts := util.ParseInputDelimiter(p[1:len(p)-1], []byte(","))
			value := uint64(0)
			for _, p := range buttonParts {
				b := util.Btoi(p)
				value |= 1 << b
			}
			buttons[i] = value
		}

		parts = util.ParseInputDelimiter(line[bytes.IndexByte(line, '{')+1:len(line)-1], []byte(","))
		joltages := make([]int, len(parts))
		for i, p := range parts {
			joltages[i] = util.Btoi(p)
		}

		machines[i] = machine{
			lights:   lights,
			buttons:  buttons,
			joltages: joltages,
		}
	}
	return machines
}

func Run(input []byte) (util.Result, error) {
	start := time.Now()

	machines := parseInput(input)

	parse := time.Now()

	part1 := 0
	part2 := -1

	for _, machine := range machines {
		minPress := 999999999
		for perm := uint64(1); perm < (1 << len(machine.buttons)); perm++ {
			lights := uint64(0)
			for b := 0; b < len(machine.buttons); b++ {
				if perm&(1<<b) > 0 {
					lights ^= machine.buttons[b]
				}
			}
			if lights == machine.lights {
				minPress = min(minPress, bits.OnesCount64(perm))
			}
		}
		part1 += minPress
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
