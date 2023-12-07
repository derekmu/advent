package day07

import (
	"advent/util"
	"bytes"
	"cmp"
	"slices"
	"time"
)

type hand struct {
	cards  []byte
	bid    int
	class1 byte
	class2 byte
}

var cardMap = map[byte]byte{
	'A': 12,
	'K': 11,
	'Q': 10,
	'J': 9,
	'T': 8,
	'9': 7,
	'8': 6,
	'7': 5,
	'6': 4,
	'5': 3,
	'4': 2,
	'3': 1,
	'2': 0,
}

func Run(input []byte) error {
	start := time.Now()

	hands := make([]*hand, 0, 1000)
	lines := util.ParseInputLines(input)
	counts := [13]byte{}
	for _, line := range lines {
		cards, bid, _ := bytes.Cut(line, []byte(" "))
		counts = [13]byte{}
		class1 := byte(0)
		hasTrio, duoCount := false, 0
		for i := 0; i < 5; i++ {
			cards[i] = cardMap[cards[i]]
			counts[cards[i]]++
			if counts[cards[i]] == 5 {
				class1 = 10
			} else if counts[cards[i]] == 4 {
				class1 = 9
				hasTrio = false
			} else if counts[cards[i]] == 3 {
				hasTrio = true
				duoCount--
			} else if counts[cards[i]] == 2 {
				duoCount++
			}
		}
		if class1 == 0 {
			if hasTrio {
				if duoCount == 1 {
					class1 = 8
				} else {
					class1 = 7
				}
			} else if duoCount == 2 {
				class1 = 6
			} else if duoCount == 1 {
				class1 = 5
			}
		}
		hands = append(hands, &hand{cards: cards, bid: util.Btoi(bid), class1: class1})
	}

	parse := time.Now()

	slices.SortFunc(hands, func(a, b *hand) int {
		c := cmp.Compare(a.class1, b.class1)
		if c != 0 {
			return c
		}
		for i := 0; i < 5; i++ {
			c = cmp.Compare(a.cards[i], b.cards[i])
			if c != 0 {
				return c
			}
		}
		return 0
	})
	part1 := 0
	for i, h := range hands {
		part1 += h.bid * (i + 1)
	}

	part2 := -1

	end := time.Now()

	util.PrintResults(part1, part2, start, parse, end)

	return nil
}
