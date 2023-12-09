package day07

import (
	"bytes"
	"testing"
)

var (
	sampleInput = []byte(`32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
`)
	sampleHands = []*hand{
		{[]byte{1, 0, 8, 1, 11}, 765, onePair, onePair},
		{[]byte{8, 3, 3, 9, 3}, 684, threeKind, fourKind},
		{[]byte{11, 11, 4, 5, 5}, 28, twoPair, twoPair},
		{[]byte{11, 8, 9, 9, 8}, 220, twoPair, fourKind},
		{[]byte{10, 10, 10, 9, 12}, 483, threeKind, fourKind},
	}
)

func TestParseInput(t *testing.T) {
	hands := parseInput(sampleInput)

	if len(sampleHands) != len(hands) {
		t.Fatal("incorrect size of node map")
	}
	for i, h := range hands {
		sh := sampleHands[i]
		if !bytes.Equal(h.cards, sh.cards) {
			t.Fatal("incorrect cards")
		}
		if h.bid != sh.bid {
			t.Fatal("incorrect bid")
		}
		if h.class1 != sh.class1 {
			t.Fatal("incorrect class1")
		}
		if h.class2 != sh.class2 {
			t.Fatal("incorrect class2")
		}
	}
}

func TestRunSample(t *testing.T) {
	result, err := Run(sampleInput)

	if err != nil {
		t.Fatal("unexpected error")
	}
	if result.Part1 != 6440 {
		t.Fatal("incorrect part 1")
	}
	if result.Part2 != 5905 {
		t.Fatal("incorrect part 2")
	}
}

func TestRun(t *testing.T) {
	result, err := Run(Input)

	if err != nil {
		t.Fatal("unexpected error")
	}
	if result.Part1 != 253866470 {
		t.Fatal("incorrect part 1")
	}
	if result.Part2 != 254494947 {
		t.Fatal("incorrect part 2")
	}
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result, err := Run(Input)
		_, _ = result, err
	}
}
