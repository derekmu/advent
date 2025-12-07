package day05

import (
	"advent/util"
	"bytes"
	_ "embed"
	"slices"
	"sort"
	"time"
)

//go:embed input.txt
var Input []byte
var Problem = util.Problem{Year: "2023", Day: "05", Runner: Run, Input: Input}

type ranger struct {
	first int
	last  int
}

type rangeMap struct {
	destFirst   int
	sourceFirst int
	sourceLast  int
}

type mapper []rangeMap

func parseInput(input []byte) ([]int, []mapper) {
	lines := util.ParseInputLines(input)
	line, _ := bytes.CutPrefix(lines[0], []byte("seeds: "))
	seeds := make([]int, 0, 20)
	for len(line) > 0 {
		num, line2, _ := bytes.Cut(line, []byte(" "))
		seeds = append(seeds, util.Btoi(num))
		line = line2
	}
	lines = lines[2:]
	mappers := make([]mapper, 0, 7)
	for len(lines) > 0 {
		endI := slices.IndexFunc(lines, func(line []byte) bool {
			return len(line) == 0
		})
		mapperLines := lines
		if endI >= 0 {
			mapperLines = lines[:endI]
			lines = lines[endI+1:]
		} else {
			lines = lines[:0]
		}
		rm := parseMapper(mapperLines)
		mappers = append(mappers, rm)
	}
	return seeds, mappers
}

func Run(input []byte) (util.Result, error) {
	start := time.Now()

	seeds, mappers := parseInput(input)

	parse := time.Now()

	part1 := -1
	for _, v := range seeds {
		for _, m := range mappers {
			mi := sort.Search(len(m), func(mi int) bool {
				return v <= m[mi].sourceLast
			})
			if mi < len(m) && v >= m[mi].sourceFirst {
				v = m[mi].destFirst + (v - m[mi].sourceFirst)
			}
		}
		if part1 == -1 || v < part1 {
			part1 = v
		}
	}

	rangers := make([]ranger, 0, 46)
	newRangers := make([]ranger, 0, 46)
	part2 := -1
	for i := 0; i < len(seeds); i += 2 {
		rangers = append(rangers, ranger{first: seeds[i], last: seeds[i] + seeds[i+1] - 1})
		for _, m := range mappers {
			for _, r1 := range rangers {
				newRangers = overlap(r1, m, newRangers)
			}
			rangers, newRangers = newRangers, rangers
			newRangers = newRangers[:0]
		}
		for _, r := range rangers {
			if part2 == -1 || r.first < part2 {
				part2 = r.first
			}
		}
		rangers = rangers[:0]
		newRangers = newRangers[:0]
	}

	end := time.Now()

	return util.Result{
		Part1:     part1,
		Part2:     part2,
		StartTime: start,
		ParseTime: parse,
		EndTime:   end,
	}, nil
}

func overlap(r ranger, m mapper, result []ranger) []ranger {
	v := r.first
	mi := sort.Search(len(m), func(mi int) bool {
		return v <= m[mi].sourceLast
	})
	for v <= r.last && mi < len(m) {
		if m[mi].sourceFirst > r.last {
			// range is after the input
			break
		}
		if v > m[mi].sourceLast {
			// range is before the input
			mi++
			continue
		}
		if v < m[mi].sourceFirst {
			// unmapped values, map to same
			result = append(result, ranger{
				first: v,
				last:  m[mi].sourceFirst - 1,
			})
			v = m[mi].sourceFirst
		}
		last := min(r.last, m[mi].sourceLast)
		// overlap, map to destination
		result = append(result, ranger{
			first: m[mi].destFirst + v - m[mi].sourceFirst,
			last:  m[mi].destFirst + last - m[mi].sourceFirst,
		})
		v = last + 1
		mi++
	}
	if v <= r.last {
		// no more ranges to consider for overlaps, map to same
		result = append(result, ranger{
			first: v,
			last:  r.last,
		})
	}
	return result
}

func parseMapper(lines [][]byte) mapper {
	lines = lines[1:]
	m := make(mapper, 0, len(lines))
	var dest, source, length []byte
	for _, line := range lines {
		dest, line, _ = bytes.Cut(line, []byte(" "))
		source, length, _ = bytes.Cut(line, []byte(" "))
		sourceFirst := util.Btoi(source)
		m = append(m, rangeMap{
			destFirst:   util.Btoi(dest),
			sourceFirst: sourceFirst,
			sourceLast:  sourceFirst + util.Btoi(length) - 1,
		})
	}
	sort.Slice(m, func(a int, b int) bool {
		return m[a].sourceFirst < m[b].sourceFirst
	})
	return m
}
