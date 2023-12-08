package day01

import (
	"advent/util"
	_ "embed"
	"time"
)

//go:embed input.txt
var Input []byte

func Run(input []byte) (*util.Result, error) {
	start := time.Now()

	lines := util.ParseInputLines(input)

	parse := time.Now()

	maxCalories := [3]int{0, 0, 0}
	calories := 0

	for _, line := range lines {
		if len(line) == 0 {
			for i := 0; i < len(maxCalories); i++ {
				if calories > maxCalories[i] {
					for j := len(maxCalories) - 1; j > i; j-- {
						maxCalories[j] = maxCalories[j-1]
					}
					maxCalories[i] = calories
					break
				}
			}
			calories = 0
		} else {
			calories += util.Btoi(line)
		}
	}

	part1 := maxCalories[0]
	part2 := maxCalories[0] + maxCalories[1] + maxCalories[2]

	end := time.Now()

	return &util.Result{
		Part1:     part1,
		Part2:     part2,
		StartTime: start,
		ParseTime: parse,
		EndTime:   end,
	}, nil
}
