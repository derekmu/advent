package day07

import (
	"advent/util"
	"bytes"
	"cmp"
	_ "embed"
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

const (
	fiveKind  = 10
	fourKind  = 9
	fullHouse = 8
	threeKind = 7
	twoPair   = 6
	onePair   = 5
	bupkis    = 4
)

//go:embed input.txt
var Input []byte

func parseInput(input []byte) (hands []*hand) {
	hands = make([]*hand, 0, 1000)
	lines := util.ParseInputLines(input)
	counts := [13]byte{}
	for _, line := range lines {
		cardsB, bid, _ := bytes.Cut(line, []byte(" "))
		cards := make([]byte, 5)
		counts = [13]byte{}
		class1 := byte(0)
		hasTrio, duoCount := false, 0
		for i := 0; i < 5; i++ {
			cards[i] = cardMap[cardsB[i]]
			counts[cards[i]]++
			if counts[cards[i]] == 5 {
				class1 = fiveKind
			} else if counts[cards[i]] == 4 {
				class1 = fourKind
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
					class1 = fullHouse
				} else {
					class1 = threeKind
				}
			} else if duoCount == 2 {
				class1 = twoPair
			} else if duoCount == 1 {
				class1 = onePair
			} else {
				class1 = bupkis
			}
		}
		jokerCount := counts[9]
		class2 := class1
		if jokerCount == 4 {
			class2 = fiveKind
		} else if jokerCount == 3 {
			if duoCount == 1 {
				class2 = fiveKind
			} else {
				class2 = fourKind
			}
		} else if jokerCount == 2 {
			if hasTrio {
				class2 = fiveKind
			} else if duoCount == 2 {
				class2 = fourKind
			} else {
				class2 = threeKind
			}
		} else if jokerCount == 1 {
			if class2 == fourKind {
				class2 = fiveKind
			} else if hasTrio {
				class2 = fourKind
			} else if duoCount == 2 {
				class2 = fullHouse
			} else if duoCount == 1 {
				class2 = threeKind
			} else {
				class2 = onePair
			}
		}
		hands = append(hands, &hand{cards: cards, bid: util.Btoi(bid), class1: class1, class2: class2})
	}
	return hands
}

func Run(input []byte) (*util.Result, error) {
	start := time.Now()

	hands := parseInput(input)

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

	slices.SortFunc(hands, func(a, b *hand) int {
		c := cmp.Compare(a.class2, b.class2)
		if c != 0 {
			return c
		}
		for i := 0; i < 5; i++ {
			if a.cards[i] == 9 && b.cards[i] == 9 {
				continue
			} else if a.cards[i] == 9 {
				return -1
			} else if b.cards[i] == 9 {
				return 1
			}
			c = cmp.Compare(a.cards[i], b.cards[i])
			if c != 0 {
				return c
			}
		}
		return 0
	})
	part2 := 0
	for i, h := range hands {
		part2 += h.bid * (i + 1)
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
