package day12

import (
	"advent/util"
	_ "embed"
	"time"
)

//go:embed input.txt
var Input []byte
var Problem = util.Problem{Year: "2025", Day: "12", Runner: Run, Input: Input}

type region struct {
	width, height int
	pieces        []int
}

func parseInput(input []byte) (regions []region) {
	lines := util.ParseInputLines(input)
	lines = lines[30:]
	regions = make([]region, len(lines))
	for i, line := range lines {
		parts := util.ParseInputDelimiter(line, []byte(": "))
		sizeParts := util.ParseInputDelimiter(parts[0], []byte("x"))
		piecesParts := util.ParseInputDelimiter(parts[1], []byte(" "))
		pieces := make([]int, len(piecesParts))
		for i, p := range piecesParts {
			pieces[i] = util.Btoi(p)
		}
		regions[i] = region{
			width:  util.Btoi(sizeParts[0]),
			height: util.Btoi(sizeParts[1]),
			pieces: pieces,
		}
	}
	return regions
}

func Run(input []byte) (util.Result, error) {
	start := time.Now()

	regions := parseInput(input)

	parse := time.Now()

	part1 := 0
	part2 := -1

	for _, r := range regions {
		sumPieces := 0
		for _, c := range r.pieces {
			sumPieces += c
		}
		rowSize := r.width / 3
		colSize := r.height / 3
		if rowSize*colSize >= sumPieces {
			part1++
		}
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
