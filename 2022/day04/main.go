package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	containedCount := 0
	overlapCount := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		nums1 := parseNums(parts[0])
		nums2 := parseNums(parts[1])
		if nums1[0] <= nums2[0] && nums1[1] >= nums2[1] {
			containedCount++
		} else if nums2[0] <= nums1[0] && nums2[1] >= nums1[1] {
			containedCount++
		}
		if nums1[0] <= nums2[1] && nums1[1] >= nums2[0] {
			overlapCount++
		}
	}
	log.Printf("The number of pairs with a range fully containing the other is %d.", containedCount)
	log.Printf("The number of pairs with an overlap is %d.", overlapCount)
}

func parseNums(s string) [2]int {
	parts := strings.Split(s, "-")
	nums := [2]int{}
	for i := 0; i < len(nums); i++ {
		v, err := strconv.Atoi(parts[i])
		if err != nil {
			log.Panic(err)
		}
		nums[i] = v
	}
	return nums
}
