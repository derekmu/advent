package main

import (
	"advent/util"
	"bytes"
	"slices"
	"time"
)

type card struct {
	winners []int
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
			winners: parseNumbers(winners),
			numbers: parseNumbers(numbers),
		}
		cards = append(cards, c)
	}

	parse := time.Now()

	part1 := 0
	for _, c := range cards {
		points := 0
		wini := 0
		numi := 0
		for wini < len(c.winners) && numi < len(c.numbers) {
			if c.winners[wini] == c.numbers[numi] {
				c.matches++
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
				wini++
				numi++
			} else if c.winners[wini] < c.numbers[numi] {
				wini++
			} else {
				numi++
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

func parseNumbers(bytes []byte) []int {
	result := make([]int, 0, (len(bytes)+1)/3)
	for i := 0; i < len(bytes); i += 3 {
		num := bytes[i : i+2]
		if num[0] == ' ' {
			num = num[1:]
		}
		result = append(result, util.Btoi(num))
	}
	slices.Sort(result)
	return result
}
