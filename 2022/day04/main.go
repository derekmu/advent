package day04

import (
	"advent/util"
	"bytes"
	_ "embed"
	"time"
)

//go:embed input.txt
var Input []byte
var Problem = util.Problem{Year: "2022", Day: "04", Runner: Run, Input: Input}

func Run(input []byte) (util.Result, error) {
	start := time.Now()

	lines := util.ParseInputLines(input)

	parse := time.Now()

	part1 := 0
	part2 := 0
	for _, line := range lines {
		parts := bytes.Split(line, []byte(","))
		nums1 := parseNums(parts[0])
		nums2 := parseNums(parts[1])
		if nums1[0] <= nums2[0] && nums1[1] >= nums2[1] {
			part1++
		} else if nums2[0] <= nums1[0] && nums2[1] >= nums1[1] {
			part1++
		}
		if nums1[0] <= nums2[1] && nums1[1] >= nums2[0] {
			part2++
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

func parseNums(s []byte) [2]int {
	parts := bytes.Split(s, []byte("-"))
	nums := [2]int{}
	for i := 0; i < len(nums); i++ {
		nums[i] = util.Btoi(parts[i])
	}
	return nums
}
