package day08

import (
	"bytes"
	"testing"
)

var (
	sampleInput = []byte(`RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)
`)
	sampleInput2 = []byte(`LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)
`)
	sampleNodeMap = map[uint32]*node{
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
	moves, nodeMap := parseInput(sampleInput)

	if !bytes.Equal(moves, []byte("RL")) {
		t.Fatal("incorrect moves parsed")
	}
	if len(sampleNodeMap) != len(nodeMap) {
		t.Fatal("incorrect size of node map")
	}
	for k, n := range nodeMap {
		sn := sampleNodeMap[k]
		if sn.left != n.left {
			t.Fatal("incorrect left")
		}
		if sn.right != n.right {
			t.Fatal("incorrect right")
		}
	}
}

func TestRunSample(t *testing.T) {
	result, err := Run(sampleInput)

	if err != nil {
		t.Fatal("unexpected error")
	}
	if result.Part1 != 2 {
		t.Fatal("incorrect part 1")
	}
	if result.Part2 != 2 {
		t.Fatal("incorrect part 2")
	}
}

func TestRunSample2(t *testing.T) {
	result, err := Run(sampleInput2)

	if err != nil {
		t.Fatal("unexpected error")
	}
	if result.Part1 != -1 {
		t.Fatal("incorrect part 1")
	}
	if result.Part2 != 6 {
		t.Fatal("incorrect part 2")
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
		t.Fatal("incorrect part 2")
	}
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result, err := Run(Input)
		_, _ = result, err
	}
}
