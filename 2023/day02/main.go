package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
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
	log.SetFlags(0)
	file, err := os.Open("input.txt")
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	start := time.Now()

	var games []*game
	for gameId := 1; scanner.Scan(); gameId++ {
		line := scanner.Text()
		g := &game{gameId: gameId}
		games = append(games, g)
		line = line[strings.Index(line, ": ")+2:]
		for len(line) > 0 {
			li := strings.Index(line, "; ")
			if li < 0 {
				li = len(line)
			}
			setLine := line[:li]
			line = line[min(len(line), li+2):]
			s := &set{}
			g.sets = append(g.sets, s)
			for len(setLine) > 0 {
				ci := strings.Index(setLine, ", ")
				if ci < 0 {
					ci = len(setLine)
				}
				cubeLine := setLine[:ci]
				setLine = setLine[min(len(setLine), ci+2):]
				si := strings.Index(cubeLine, " ")
				count, err := strconv.Atoi(cubeLine[:si])
				if err != nil {
					log.Panic(err)
				}
				switch cubeLine[si+1:] {
				case "red":
					s.redCount = count
				case "green":
					s.greenCount = count
				case "blue":
					s.blueCount = count
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

	log.Printf("The sum of the game ids possible is %d", part1)
	log.Printf("The sum of the power of the sets is %d", part2)
	log.Printf("Parse time was %v", parse.Sub(start))
	log.Printf("Execute time was %v", end.Sub(parse))
}
