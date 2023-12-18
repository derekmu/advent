package day15

import (
	"advent/util"
	"bytes"
	_ "embed"
	"time"
)

var Problem = util.Problem{Year: "2023", Day: "15", Runner: Run, Input: Input}

type lens struct {
	label uint64
	focal byte
}

//go:embed input.txt
var Input []byte

func Run(input []byte) (*util.Result, error) {
	start := time.Now()

	input = input[:len(input)-1] // strip new line

	parse := time.Now()

	part1 := 0
	boxLenses := make([][]*lens, 256)
	var line []byte
	ok := true
	for ok {
		line, input, ok = bytes.Cut(input, []byte(","))
		part1 += int(hash(line))
		if line[len(line)-1] == '-' {
			label := line[:len(line)-1]
			l2 := smash(label)
			box := hash(label)
			lenses := boxLenses[box]
			for i, l := range lenses {
				if l.label == l2 {
					copy(lenses[i:], lenses[i+1:])
					boxLenses[box] = lenses[:len(lenses)-1]
					break
				}
			}
		} else {
			label, focalChar, _ := bytes.Cut(line, []byte("="))
			l2 := smash(label)
			focal := focalChar[0] - '0'
			found := false
			box := hash(label)
			lenses := boxLenses[box]
			for _, l := range lenses {
				if l.label == l2 {
					l.focal = focal
					found = true
					break
				}
			}
			if !found {
				if boxLenses[box] == nil {
					boxLenses[box] = make([]*lens, 0, 7)
				}
				boxLenses[box] = append(boxLenses[box], &lens{
					label: l2,
					focal: focal,
				})
			}
		}
	}
	part2 := 0
	maxCap := 0
	for bi, lenses := range boxLenses {
		for li, l := range lenses {
			part2 += (bi + 1) * (li + 1) * int(l.focal)
		}
		if cap(lenses) > maxCap {
			maxCap = cap(lenses)
		}
	}

	end := time.Now()

	return &util.Result{
		Part1:     part1,
		Part2:     part2,
		StartTime: start,
		ParseTime: parse,
		EndTime:   end,
	}, nil
}

func hash(line []byte) uint8 {
	h := uint8(0)
	for _, c := range line {
		h += c
		h *= 17
	}
	return h
}

func smash(line []byte) uint64 {
	h := uint64(0)
	for i, c := range line {
		h |= uint64(c) << (i * 8)
	}
	return h
}
