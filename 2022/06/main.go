package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)
	file, err := os.Open("input.txt")
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		chars4 := map[uint8]int{}
		fin4 := false
		chars14 := map[uint8]int{}
		fin14 := false
		for i := 0; i < len(line) && (!fin4 || !fin14); i++ {
			if !fin4 {
				fin4 = updateCharCountMap(chars4, 4, line, i)
			}
			if !fin14 {
				fin14 = updateCharCountMap(chars14, 14, line, i)
			}
		}
	}
}

func updateCharCountMap(charCountMap1 map[uint8]int, size int, line string, i int) bool {
	c := line[i]
	v, _ := charCountMap1[c]
	charCountMap1[c] = v + 1
	if i >= size {
		c1 := line[i-size]
		v1, _ := charCountMap1[c1]
		if v1 == 1 {
			delete(charCountMap1, c1)
		} else {
			charCountMap1[c1] = v1 - 1
		}
	}
	if len(charCountMap1) == size {
		log.Printf("The number of characters to process before the first start of packet of size %d is %d", size, i+1)
		return true
	}
	return false
}
