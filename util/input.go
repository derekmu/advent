package util

import (
	"bytes"
	"log"
	"os"
)

var newline = []byte{'\n'}

func ReadInput() []byte {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Panic(err)
	}
	return input
}

func ParseInputLines(input []byte) [][]byte {
	count := bytes.Count(input, newline)
	result := make([][]byte, 0, count)
	var line []byte
	var found bool
	for row := 0; ; row++ {
		line, input, found = bytes.Cut(input, newline)
		if !found {
			break
		}
		result = append(result, line)
	}
	return result
}
