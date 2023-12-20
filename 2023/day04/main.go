package day04

import (
	"advent/util"
	"bytes"
	_ "embed"
	"slices"
	"time"
)

//go:embed input.txt
var Input []byte
var Problem = util.Problem{Year: "2023", Day: "04", Runner: Run, Input: Input}

type card struct {
	matches int
	copies  int
}

func parseInput(input []byte) []*card {
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
	return cards
}

func Run(input []byte) (*util.Result, error) {
	start := time.Now()

	cards := parseInput(input)

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

	return &util.Result{
		Part1:     part1,
		Part2:     part2,
		StartTime: start,
		ParseTime: parse,
		EndTime:   end,
	}, nil
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
