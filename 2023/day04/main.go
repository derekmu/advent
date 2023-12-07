package day04

import (
	"advent/util"
	"bytes"
	"slices"
	"time"
)

type card struct {
	matches int
	copies  int
}

func Run(input []byte) error {
	start := time.Now()

	lines := util.ParseInputLines(input)
	cards := make([]*card, 0, len(lines))
	for _, line := range lines {
		_, line, _ = bytes.Cut(line, []byte(": "))
		winners, numbers, _ := bytes.Cut(line, []byte(" | "))
		winers := parseNumbers(winners)
		matches := 0
		for i := 0; i < len(numbers); i += 3 {
			num := uint16(numbers[i])<<8 + uint16(numbers[i+1])
			_, found := slices.BinarySearch(winers, num)
			if found {
				matches++
			}
		}
		c := &card{
			matches: matches,
			copies:  1,
		}
		cards = append(cards, c)
	}

	parse := time.Now()

	part1 := 0
	for _, c := range cards {
		if c.matches > 0 {
			part1 += 1 << (c.matches - 1)
		}
	}

	part2 := 0
	for ci, c := range cards {
		for ci2 := ci + 1; ci2 <= ci+c.matches && ci2 < len(cards); ci2++ {
			cards[ci2].copies += c.copies
		}
	}
	for _, c := range cards {
		part2 += c.copies
	}

	end := time.Now()

	util.PrintResults(part1, part2, start, parse, end)

	return nil
}

func parseNumbers(bytes []byte) []uint16 {
	result := make([]uint16, 0, (len(bytes)+1)/3)
	for i := 0; i < len(bytes); i += 3 {
		num := uint16(bytes[i])<<8 + uint16(bytes[i+1])
		result = append(result, num)
	}
	slices.Sort(result)
	return result
}
