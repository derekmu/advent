package day07

import (
	"bytes"
	"testing"
)

var (
	testInput = []byte(`RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)
`)
	expectedNodeMap = map[uint32]*node{
		stoui("AAA"): {stoui("BBB"), stoui("CCC")},
		stoui("BBB"): {stoui("DDD"), stoui("EEE")},
		stoui("CCC"): {stoui("ZZZ"), stoui("GGG")},
		stoui("DDD"): {stoui("DDD"), stoui("DDD")},
		stoui("EEE"): {stoui("EEE"), stoui("EEE")},
		stoui("GGG"): {stoui("GGG"), stoui("GGG")},
		stoui("ZZZ"): {stoui("ZZZ"), stoui("ZZZ")},
	}
)

func TestParseInput(t *testing.T) {
	moves, nodeMap := parseInput(testInput)

	if !bytes.Equal(moves, []byte("RL")) {
		t.Fatal("incorrect moves parsed")
	}
	if len(expectedNodeMap) != len(nodeMap) {
		t.Fatal("incorrect size of node map")
	}
	for k, v := range nodeMap {
		ev := expectedNodeMap[k]
		if ev.left != v.left {
			t.Fatal("incorrect left")
		}
		if ev.right != v.right {
			t.Fatal("incorrect right")
		}
	}
}

func TestRun(t *testing.T) {
	result, err := Run(Input)

	if err != nil {
		t.Fatal("unexpected error")
	}
	if result.Part1 != 14257 {
		t.Fatal("incorrect part 1")
	}
	if result.Part2 != 16187743689077 {
		t.Fatal("incorrect part 1")
	}
}
