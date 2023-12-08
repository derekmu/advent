package day02

import (
	"advent/util"
	"bufio"
	"bytes"
	_ "embed"
	"errors"
	"time"
)

var (
	winsMap = map[string]string{
		"rock":    "scissor",
		"scissor": "paper",
		"paper":   "rock",
	}
	losesMap = map[string]string{
		"rock":    "paper",
		"scissor": "rock",
		"paper":   "scissor",
	}
	choiceMap = map[string]string{
		"A": "rock",
		"B": "paper",
		"C": "scissor",
		"X": "rock",
		"Y": "paper",
		"Z": "scissor",
	}
	scoreMap = map[string]int{
		"rock":    1,
		"paper":   2,
		"scissor": 3,
	}
)

//go:embed input.txt
var Input []byte

func Run(input []byte) (*util.Result, error) {
	start := time.Now()

	parse := time.Now()

	part1 := 0
	part2 := 0

	scanner := bufio.NewScanner(bytes.NewReader(input))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		theirChoice := choiceMap[scanner.Text()]
		if !scanner.Scan() {
			return nil, errors.New("missing second token")
		}
		me := scanner.Text()
		myChoice := choiceMap[me]
		if myChoice == theirChoice {
			part1 += 3
		} else if winsMap[myChoice] == theirChoice {
			part1 += 6
		}
		part1 += scoreMap[myChoice]

		switch me {
		case "X":
			myChoice = winsMap[theirChoice]
		case "Y":
			myChoice = theirChoice
		case "Z":
			myChoice = losesMap[theirChoice]
		}
		if myChoice == theirChoice {
			part2 += 3
		} else if winsMap[myChoice] == theirChoice {
			part2 += 6
		}
		part2 += scoreMap[myChoice]
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
