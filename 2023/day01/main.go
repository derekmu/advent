package day01

import (
	"advent/util"
	"bytes"
	_ "embed"
	"time"
)

//go:embed input.txt
var Input []byte
var Problem = util.Problem{Year: "2023", Day: "01", Runner: Run, Input: Input}

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

func Run(input []byte) (*util.Result, error) {
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

	return &util.Result{
		Part1:     part1,
		Part2:     part2,
		StartTime: start,
		ParseTime: parse,
		EndTime:   end,
	}, nil
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
