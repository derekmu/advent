package day02

import (
	"bufio"
	"bytes"
	"errors"
	"log"
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

func Run(input []byte) error {
	totalScore1 := 0
	totalScore2 := 0

	scanner := bufio.NewScanner(bytes.NewReader(input))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		theirChoice := choiceMap[scanner.Text()]
		if !scanner.Scan() {
			return errors.New("missing second token")
		}
		me := scanner.Text()
		myChoice := choiceMap[me]
		if myChoice == theirChoice {
			totalScore1 += 3
		} else if winsMap[myChoice] == theirChoice {
			totalScore1 += 6
		}
		totalScore1 += scoreMap[myChoice]

		switch me {
		case "X":
			myChoice = winsMap[theirChoice]
		case "Y":
			myChoice = theirChoice
		case "Z":
			myChoice = losesMap[theirChoice]
		}
		if myChoice == theirChoice {
			totalScore2 += 3
		} else if winsMap[myChoice] == theirChoice {
			totalScore2 += 6
		}
		totalScore2 += scoreMap[myChoice]
	}

	log.Printf("The total score for part 1 would be %d.", totalScore1)
	log.Printf("The total score for part 2 would be %d.", totalScore2)

	return nil
}
