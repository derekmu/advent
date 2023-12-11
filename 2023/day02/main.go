package day02

import (
	"advent/util"
	"bytes"
	_ "embed"
	"errors"
	"fmt"
	"time"
)

type set struct {
	red   int
	green int
	blue  int
}

type game struct {
	gameId int
	sets   []*set
}

//go:embed input.txt
var Input []byte

func parseInput(input []byte) ([]*game, error) {
	lines := util.ParseInputLines(input)
	games := make([]*game, 0, len(lines))
	var setLine []byte
	for gameId, line := range lines {
		g := &game{gameId: gameId + 1}
		games = append(games, g)
		_, line, _ = bytes.Cut(line, []byte(": "))
		for len(line) > 0 {
			setLine, line, _ = bytes.Cut(line, []byte("; "))
			s := &set{}
			g.sets = append(g.sets, s)
			for len(setLine) > 0 {
				ci := bytes.Index(setLine, []byte(", "))
				if ci < 0 {
					ci = len(setLine)
				}
				cubeLine := setLine[:ci]
				setLine = setLine[min(len(setLine), ci+2):]
				si := bytes.Index(cubeLine, []byte(" "))
				count := util.Btoi(cubeLine[:si])
				color := cubeLine[si+1:]
				if bytes.Equal(color, []byte("red")) {
					s.red = count
				} else if bytes.Equal(color, []byte("green")) {
					s.green = count
				} else if bytes.Equal(color, []byte("blue")) {
					s.blue = count
				} else {
					return nil, errors.New(fmt.Sprintf("Unknown color %s", color))
				}
			}
		}
	}
	return games, nil
}

func Run(input []byte) (*util.Result, error) {
	start := time.Now()

	games, err := parseInput(input)
	if err != nil {
		return nil, err
	}

	parse := time.Now()

	part1 := 0
	part2 := 0
	for _, g := range games {
		impossible := false
		for _, s := range g.sets {
			if s.red > 12 || s.green > 13 || s.blue > 14 {
				impossible = true
				break
			}
		}
		if !impossible {
			part1 += g.gameId
		}

		maxRed := 0
		maxGreen := 0
		maxBlue := 0
		for _, s := range g.sets {
			maxRed = max(maxRed, s.red)
			maxGreen = max(maxGreen, s.green)
			maxBlue = max(maxBlue, s.blue)
		}
		part2 += maxRed * maxGreen * maxBlue
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
