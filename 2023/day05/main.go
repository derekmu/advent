package main

import (
	"advent/util"
	"bytes"
	"log"
	"slices"
	"sort"
	"time"
)

type ranger struct {
	start  int
	length int
}

type rangeMap struct {
	dest   int
	source int
	length int
}

type mapper []rangeMap

func main() {
	input := util.ReadInput()

	start := time.Now()

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

	parse := time.Now()

	part1 := -1
	for _, v := range seeds {
		for _, m := range mappers {
			for mi := 0; mi < len(m); mi++ {
				if m[mi].source <= v && m[mi].source+m[mi].length > v {
					v = m[mi].dest + (v - m[mi].source)
					break
				}
			}
		}
		if part1 == -1 || v < part1 {
			part1 = v
		}
	}

	part1t := time.Now()

	part2 := -1
	for i := 0; i < len(seeds); i += 2 {
		rangers := make([]ranger, 0, 100)
		rangers = append(rangers, ranger{start: seeds[i], length: seeds[i+1]})
		for _, m := range mappers {
			newRangers := make([]ranger, 0, 100)
			for _, r := range rangers {
				for _, r2 := range overlap(r, m) {
					newRangers = append(newRangers, r2)
				}
			}
			rangers = newRangers
		}
		for _, r := range rangers {
			if part2 == -1 || r.start < part2 {
				part2 = r.start
			}
		}
	}

	end := time.Now()

	util.PrintResults(part1, part2, start, parse, end)
	log.Printf("Timing (part1 / part 2): %v / %v", part1t.Sub(parse), end.Sub(part1t))
}

func overlap(r ranger, m mapper) []ranger {
	result := make([]ranger, 0, len(m))
	v := r.start
	mi := 0
	for v < r.start+r.length && mi < len(m) {
		rm := m[mi]
		if rm.source >= r.start+r.length {
			// range is after the input
			break
		}
		if v >= rm.source+rm.length {
			// range is before the input
			mi++
			continue
		}
		if v < rm.source {
			// unmapped values, map to same
			result = append(result, ranger{
				start:  v,
				length: rm.source - v,
			})
			v = rm.source
		}
		end := min(r.start+r.length, rm.source+rm.length)
		// overlap, map to destination
		result = append(result, ranger{
			start:  rm.dest + (v - rm.source),
			length: end - v,
		})
		v = end
		mi++
	}
	if v < r.start+r.length {
		// no more ranges to consider for overlaps, map to same
		result = append(result, ranger{
			start:  v,
			length: r.start + r.length - v,
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
		m = append(m, rangeMap{
			dest:   util.Btoi(dest),
			source: util.Btoi(source),
			length: util.Btoi(length),
		})
	}
	sort.Slice(m, func(a int, b int) bool {
		return m[a].source < m[b].source
	})
	return m
}
