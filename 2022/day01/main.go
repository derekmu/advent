package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	maxCalories := [3]int{0, 0, 0}
	calories := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
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
			lcal, err := strconv.Atoi(line)
			if err != nil {
				log.Panic(err)
			}
			calories += lcal
		}
	}

	if err := scanner.Err(); err != nil {
		log.Panic(err)
	}

	log.Printf("The most calories that a single elf is carrying is %d.", maxCalories[0])
	log.Printf("The most calories that the the top 3 elves are carrying is %d.", maxCalories[0]+maxCalories[1]+maxCalories[2])
}
