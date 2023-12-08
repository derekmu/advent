package day01

import (
	"advent/util"
	_ "embed"
	"log"
)

//go:embed input.txt
var Input []byte

func Run(input []byte) error {
	maxCalories := [3]int{0, 0, 0}
	calories := 0

	lines := util.ParseInputLines(input)
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

	log.Printf("The most calories that a single elf is carrying is %d.", maxCalories[0])
	log.Printf("The most calories that the the top 3 elves are carrying is %d.", maxCalories[0]+maxCalories[1]+maxCalories[2])

	return nil
}
