package main

import (
	"advent/util"
	"bytes"
	"time"
)

type card struct {
	winners map[int]bool
	numbers []int
	matches int
	copies  int
}

func main() {
	input := util.ReadInput()

	start := time.Now()

	lines := util.ParseInputLines(input)
	cards := make([]*card, 0, len(lines))
	for _, line := range lines {
		_, line, _ = bytes.Cut(line, []byte(": "))
		winners, numbers, _ := bytes.Cut(line, []byte(" | "))
		c := &card{
			winners: parseWinners(winners),
			numbers: parseNumbers(numbers),
		}
		cards = append(cards, c)
	}

	parse := time.Now()

	part1 := 0
	for _, c := range cards {
		points := 0
		for _, num := range c.numbers {
			if _, ok := c.winners[num]; ok {
				c.matches++
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
			}
		}
		part1 += points
	}

	part2 := 0
	for ci, c := range cards {
		for ci2 := ci + 1; ci2 <= ci+c.matches && ci2 < len(cards); ci2++ {
			cards[ci2].copies += 1 + c.copies
		}
	}
	for _, c := range cards {
		part2 += 1 + c.copies
	}

	end := time.Now()

	util.PrintResults(part1, part2, start, parse, end)
}

func parseWinners(bytes []byte) map[int]bool {
	result := make(map[int]bool, (len(bytes)+1)/3)
	for i := 0; i < len(bytes); i += 3 {
		num := bytes[i : i+2]
		if num[0] == ' ' {
			num = num[1:]
		}
		result[util.Btoi(num)] = true
	}
	return result
}

func parseNumbers(bytes []byte) []int {
	result := make([]int, (len(bytes)+1)/3)
	for i := 0; i < len(bytes); i += 3 {
		num := bytes[i : i+2]
		if num[0] == ' ' {
			num = num[1:]
		}
		result = append(result, util.Btoi(num))
	}
	return result

}
