package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var words = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

func main() {
	log.SetFlags(0)
	file, err := os.Open("input.txt")
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	sum1 := 0
	sum2 := 0
	for scanner.Scan() {
		line := scanner.Text()
		sum1 += calibrationValue1(line)
		sum2 += calibrationValue2(line)
	}

	log.Printf("The sum of the calibration values for part 1 is %d", sum1)
	log.Printf("The sum of the calibration values for part 2 is %d", sum2)
}

func calibrationValue1(line string) int {
	firstDigitIndex := strings.IndexAny(line, "123456789")
	firstDigit := int(line[firstDigitIndex] - '0')
	lastDigitIndex := strings.LastIndexAny(line, "123456789")
	lastDigit := int(line[lastDigitIndex] - '0')
	return firstDigit*10 + lastDigit
}

func calibrationValue2(line string) int {
	firstDigitIndex := strings.IndexAny(line, "123456789")
	firstDigit := int(line[firstDigitIndex] - '0')
	lastDigitIndex := strings.LastIndexAny(line, "123456789")
	lastDigit := int(line[lastDigitIndex] - '0')

	for wi, word := range words {
		firstIndex := strings.Index(line, word)
		lastIndex := strings.LastIndex(line, word)
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
