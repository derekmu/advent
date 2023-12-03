package main

import (
	"advent/util"
	"bytes"
	"log"
	"time"
)

type set struct {
	redCount   int
	greenCount int
	blueCount  int
}

type game struct {
	gameId int
	sets   []*set
}

func main() {
	input := util.ReadInput()

	start := time.Now()

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
					s.redCount = count
				} else if bytes.Equal(color, []byte("green")) {
					s.greenCount = count
				} else if bytes.Equal(color, []byte("blue")) {
					s.blueCount = count
				} else {
					log.Panicf("Unknown color %s", color)
				}
			}
		}
	}

	parse := time.Now()

	part1 := 0
	part2 := 0
	for _, g := range games {
		impossible := false
		for _, s := range g.sets {
			if s.redCount > 12 || s.greenCount > 13 || s.blueCount > 14 {
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
			maxRed = max(maxRed, s.redCount)
			maxGreen = max(maxGreen, s.greenCount)
			maxBlue = max(maxBlue, s.blueCount)
		}
		part2 += maxRed * maxGreen * maxBlue
	}

	end := time.Now()

	util.PrintResults(part1, part2, start, parse, end)
}
