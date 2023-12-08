package day04

import (
	"advent/util"
	"bytes"
	_ "embed"
	"log"
)

//go:embed input.txt
var Input []byte

func Run(input []byte) error {
	containedCount := 0
	overlapCount := 0

	lines := util.ParseInputLines(input)
	for _, line := range lines {
		parts := bytes.Split(line, []byte(","))
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

	return nil
}

func parseNums(s []byte) [2]int {
	parts := bytes.Split(s, []byte("-"))
	nums := [2]int{}
	for i := 0; i < len(nums); i++ {
		nums[i] = util.Btoi(parts[i])
	}
	return nums
}
