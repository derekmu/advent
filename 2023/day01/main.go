package day01

import (
	"advent/util"
	"bytes"
	"time"
)

var words = [][]byte{
	[]byte("one"),
	[]byte("two"),
	[]byte("three"),
	[]byte("four"),
	[]byte("five"),
	[]byte("six"),
	[]byte("seven"),
	[]byte("eight"),
	[]byte("nine"),
}

func Run(input []byte) error {
	start := time.Now()

	lines := util.ParseInputLines(input)

	parse := time.Now()

	part1 := 0
	part2 := 0
	for _, line := range lines {
		part1 += calibrationValue1(line)
		part2 += calibrationValue2(line)
	}

	end := time.Now()

	util.PrintResults(part1, part2, start, parse, end)

	return nil
}

func calibrationValue1(line []byte) int {
	firstDigitIndex := bytes.IndexAny(line, "123456789")
	firstDigit := int(line[firstDigitIndex] - '0')
	lastDigitIndex := bytes.LastIndexAny(line, "123456789")
	lastDigit := int(line[lastDigitIndex] - '0')
	return firstDigit*10 + lastDigit
}

func calibrationValue2(line []byte) int {
	firstDigitIndex := bytes.IndexAny(line, "123456789")
	firstDigit := int(line[firstDigitIndex] - '0')
	lastDigitIndex := bytes.LastIndexAny(line, "123456789")
	lastDigit := int(line[lastDigitIndex] - '0')

	for wi, word := range words {
		firstIndex := bytes.Index(line, word)
		lastIndex := bytes.LastIndex(line, word)
		if firstIndex >= 0 && firstIndex < firstDigitIndex {
			firstDigitIndex = firstIndex
			firstDigit = wi + 1
		}
		if lastIndex >= 0 && lastIndex > lastDigitIndex {
			lastDigitIndex = lastIndex
			lastDigit = wi + 1
		}
	}
	return firstDigit*10 + lastDigit
}
