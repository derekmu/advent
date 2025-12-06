package day05

import (
	"advent/util"
	_ "embed"
	"log"
	"slices"
	"time"
)

//go:embed input.txt
var Input []byte
var Problem = util.Problem{Year: "2025", Day: "05", Runner: Run, Input: Input}

var blankLine = []byte("\n\n")
var dash = []byte{'-'}

type idRange struct {
	start, end int
}

func parseInput(input []byte) (fresh []idRange, available []int) {
	parts := util.ParseInputDelimiter(input, blankLine)
	if len(parts) != 2 {
		log.Panicln("Invalid number of input parts")
	}
	lines := util.ParseInputLines(parts[0])
	allFresh := make([]idRange, len(lines))
	for i, line := range lines {
		rangeParts := util.ParseInputDelimiter(line, dash)
		if len(rangeParts) != 2 {
			log.Panicln("Invalid number of range parts")
		}
		allFresh[i] = idRange{
			start: util.Btoi(rangeParts[0]),
			end:   util.Btoi(rangeParts[1]),
		}
	}
	slices.SortFunc(allFresh, func(a, b idRange) int {
		if a.start < b.start {
			return -1
		} else if a.start > b.start {
			return 1
		} else {
			return 0
		}
	})
	fresh = make([]idRange, 0, len(allFresh))
	fi := 0
	fresh = append(fresh, allFresh[0])
	for i := 1; i < len(allFresh); i++ {
		if fresh[fi].end >= allFresh[i].start {
			fresh[fi].end = max(fresh[fi].end, allFresh[i].end)
		} else {
			fi++
			fresh = append(fresh, allFresh[i])
		}
	}

	lines = util.ParseInputLines(parts[1])
	available = make([]int, len(lines))
	for i, line := range lines {
		available[i] = util.Btoi(line)
	}
	slices.Sort(available)
	return fresh, available
}

func Run(input []byte) (*util.Result, error) {
	start := time.Now()

	fresh, available := parseInput(input)
	_, _ = fresh, available

	parse := time.Now()

	part1 := 0
	part2 := 0

	fi := 0
out:
	for _, id := range available {
		for id > fresh[fi].end {
			fi++
			if fi >= len(fresh) {
				break out
			}
		}
		if id >= fresh[fi].start && id <= fresh[fi].end {
			part1++
		}
	}

	for _, idRange := range fresh {
		part2 += idRange.end - idRange.start + 1
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
