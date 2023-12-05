package main

import (
	"advent/util"
	"bytes"
	"slices"
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

type mapper struct {
	name string
	maps []rangeMap
}

func main() {
	input := util.ReadInput()

	start := time.Now()

	lines := util.ParseInputLines(input)
	lineI := 0
	line, _ := bytes.CutPrefix(lines[lineI], []byte("seeds: "))
	var num []byte
	seeds := make([]int, 0, 20)
	for len(line) > 0 {
		num, line, _ = bytes.Cut(line, []byte(" "))
		seeds = append(seeds, util.Btoi(num))
	}
	lineI += 2
	mappers := make([]mapper, 0, 7)
	var rm mapper
	for lineI < len(lines) {
		rm, lineI = parseMapper(lines, lineI)
		mappers = append(mappers, rm)
	}

	parse := time.Now()

	part1 := -1
	for _, v := range seeds {
		for _, m := range mappers {
			v = translate(v, m)
		}
		if part1 == -1 || v < part1 {
			part1 = v
		}
	}

	part2 := -1
	for i := 0; i < len(seeds); i += 2 {
		rangers := []ranger{
			{start: seeds[i], length: seeds[i+1]},
		}
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
}

func translate(v int, m mapper) int {
	for _, rm := range m.maps {
		if rm.source <= v && rm.source+rm.length > v {
			return rm.dest + (v - rm.source)
		}
	}
	return v
}

func overlap(r ranger, m mapper) []ranger {
	result := make([]ranger, 0, len(m.maps))
	v := r.start
	mi := 0
	for v < r.start+r.length && mi < len(m.maps) {
		rm := m.maps[mi]
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

func parseMapper(lines [][]byte, li int) (mapper, int) {
	name, _ := bytes.CutSuffix(lines[li], []byte(" map:"))
	m := mapper{
		name: string(name),
		maps: make([]rangeMap, 0, 30),
	}
	li++
	var dest, source, length []byte
	for li < len(lines) {
		line := lines[li]
		li++
		if len(line) == 0 {
			return m, li
		}
		dest, line, _ = bytes.Cut(line, []byte(" "))
		source, length, _ = bytes.Cut(line, []byte(" "))
		m.maps = append(m.maps, rangeMap{
			dest:   util.Btoi(dest),
			source: util.Btoi(source),
			length: util.Btoi(length),
		})
		slices.SortFunc(m.maps, func(a rangeMap, b rangeMap) int {
			if a.source < b.source {
				return -1
			} else if b.source < a.source {
				return 1
			}
			return 0
		})
	}
	return m, li
}
