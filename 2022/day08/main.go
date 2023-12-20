package day08

import (
	"advent/util"
	_ "embed"
	"time"
)

//go:embed input.txt
var Input []byte
var Problem = util.Problem{Year: "2022", Day: "08", Runner: Run, Input: Input}

func Run(input []byte) (*util.Result, error) {
	start := time.Now()

	var viz [][]bool
	trees := util.ParseInputLines(input)
	for _, line := range trees {
		vizRow := make([]bool, len(line))
		viz = append(viz, vizRow)
		maxHeight := uint8('0' - 1)
		for i := 0; i < len(line); i++ {
			if line[i] > maxHeight {
				vizRow[i] = true
				maxHeight = line[i]
			}
		}
		maxHeight = uint8('0' - 1)
		for col := len(line) - 1; col >= 0; col-- {
			if line[col] > maxHeight {
				vizRow[col] = true
				maxHeight = line[col]
			}
		}
	}

	parse := time.Now()

	for col := 0; col < len(trees[0]); col++ {
		maxHeight := uint8('0' - 1)
		for row := 0; row < len(trees); row++ {
			if trees[row][col] > maxHeight {
				viz[row][col] = true
				maxHeight = trees[row][col]
			}
		}
		maxHeight = uint8('0' - 1)
		for row := len(trees) - 1; row >= 0; row-- {
			if trees[row][col] > maxHeight {
				viz[row][col] = true
				maxHeight = trees[row][col]
			}
		}
	}

	part1 := 0
	part2 := 0
	for row := 0; row < len(viz); row++ {
		for col := 0; col < len(viz[row]); col++ {
			if viz[row][col] {
				part1++
			}

			north := 0
			if row > 0 {
				north++
				for ro := row - 1; ro > 0 && trees[ro][col] < trees[row][col]; ro-- {
					north++
				}
			}
			south := 0
			if row < len(trees)-1 {
				south++
				for ro := row + 1; ro < len(trees)-1 && trees[ro][col] < trees[row][col]; ro++ {
					south++
				}
			}
			east := 0
			if col < len(trees[row])-1 {
				east++
				for co := col + 1; co < len(trees[row])-1 && trees[row][co] < trees[row][col]; co++ {
					east++
				}
			}
			west := 0
			if col > 0 {
				west++
				for co := col - 1; co > 0 && trees[row][co] < trees[row][col]; co-- {
					west++
				}
			}

			if north*south*east*west > part2 {
				part2 = north * south * east * west
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
