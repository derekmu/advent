package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
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

	var games []*game
	for gameId := 1; scanner.Scan(); gameId++ {
		line := scanner.Text()
		g := &game{gameId: gameId}
		games = append(games, g)
		line = strings.Split(line, ":")[1]
		setLines := strings.Split(line, ";")
		for _, setLine := range setLines {
			s := &set{}
			g.sets = append(g.sets, s)
			cubeLines := strings.Split(setLine, ",")
			for _, cubeLine := range cubeLines {
				parts := strings.Split(cubeLine, " ")
				count, err := strconv.Atoi(parts[1])
				if err != nil {
					log.Panic(err)
				}
				switch parts[2] {
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

	log.Printf("The sum of the game ids possible is %d", part1)
	log.Printf("The sum of the power of the sets is %d", part2)
}
