package day2502

import (
	"advent/util"
	_ "embed"
	"time"
)

//go:embed input.txt
var Input []byte
var Problem = util.Problem{Year: "2025", Day: "02", Runner: Run, Input: Input}

type idRange struct {
	startInt int
	endInt   int
}

var comma = []byte{','}
var dash = []byte{'-'}

func parseInput(input []byte) (ranges []idRange) {
	rangesBytes := util.ParseInputDelimiter(input, comma)
	ranges = make([]idRange, len(rangesBytes))
	for i, rangeBytes := range rangesBytes {
		rangeParts := util.ParseInputDelimiter(rangeBytes, dash)
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

func Run(input []byte) (*util.Result, error) {
	start := time.Now()

	ranges := parseInput(input)

	parse := time.Now()

	part1 := 0
	part2 := 0

	for _, r := range ranges {
		for id := r.startInt; id <= r.endInt; id++ {
			p1, p2 := invalidTest(id)
			if p1 {
				part1 += id
			}
			if p2 {
				part2 += id
			}
		}
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

func invalidTest(id int) (bool, bool) {
	l := digits(id)
	for repeats := 2; repeats <= l; repeats++ {
		if l%repeats != 0 {
			continue
		}
		size := l / repeats
		valid := false
		for j := 0; j < size; j++ {
			vj := (id / pow(10, j)) % 10
			for k := j + size; k < l; k += size {
				vk := (id / pow(10, k)) % 10
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

func digits(id int) int {
	l := 0
	for tid := id; tid > 0; tid /= 10 {
		l++
	}
	return l
}

func pow(n, m int) int {
	if m == 0 {
		return 1
	}
	if m == 1 {
		return n
	}
	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}
