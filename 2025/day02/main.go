package day2502

import (
	"advent/util"
	_ "embed"
	"sync"
	"time"
)

//go:embed input.txt
var Input []byte
var Problem = util.Problem{Year: "2025", Day: "02", Runner: Run, Input: Input}

type idRange struct {
	startInt int
	endInt   int
}

type result struct {
	part1, part2 int
}

func parseInput(input []byte) (ranges []idRange) {
	rangesBytes := util.ParseInputDelimiter(input, []byte(","))
	ranges = make([]idRange, len(rangesBytes))
	for i, rangeBytes := range rangesBytes {
		rangeParts := util.ParseInputDelimiter(rangeBytes, []byte("-"))
		if len(rangeParts) != 2 {
			panic("Invalid input")
		}
		ranges[i] = idRange{
			startInt: util.Btoi(rangeParts[0]),
			endInt:   util.Btoi(rangeParts[1]),
		}
	}
	return ranges
}

func Run(input []byte) (util.Result, error) {
	start := time.Now()

	ranges := parseInput(input)

	parse := time.Now()

	part1 := 0
	part2 := 0

	var wg sync.WaitGroup
	results := make([]result, len(ranges))
	for i, r := range ranges {
		wg.Add(1)
		go func(i int, r idRange) {
			defer wg.Done()
			for id := r.startInt; id <= r.endInt; id++ {
				p1, p2 := invalidTest(id)
				if p1 {
					results[i].part1 += id
				}
				if p2 {
					results[i].part2 += id
				}
			}
		}(i, r)
	}
	wg.Wait()
	for _, r := range results {
		part1 += r.part1
		part2 += r.part2
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

func invalidTest(id int) (bool, bool) {
	l := util.Digits(id)
	for repeats := 2; repeats <= l; repeats++ {
		if l%repeats != 0 {
			continue
		}
		size := l / repeats
		valid := false
		for j := 0; j < size; j++ {
			vj := (id / util.Pow(10, j)) % 10
			for k := j + size; k < l; k += size {
				vk := (id / util.Pow(10, k)) % 10
				if vj != vk {
					valid = true
					break
				}
			}
		}
		if !valid {
			return repeats == 2, true
		}
	}
	return false, false
}
