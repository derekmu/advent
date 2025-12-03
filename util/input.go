package util

import (
	"bytes"
)

var newline = []byte{'\n'}

func ParseInputLines(input []byte) [][]byte {
	return ParseInputDelimiter(input, newline)
}

func ParseInputDelimiter(input []byte, delimiter []byte) [][]byte {
	if len(input) > 0 && input[len(input)-1] == '\n' {
		input = input[:len(input)-1]
	}
	count := bytes.Count(input, delimiter)
	if len(input) > 0 && bytes.Compare(input[len(input)-len(delimiter):], delimiter) == 0 {
		count++
	}
	result := make([][]byte, 0, count)
	var line []byte
	var found bool
	for row := 0; ; row++ {
		line, input, found = bytes.Cut(input, delimiter)
		if found || len(line) > 0 {
			result = append(result, line)
		}
		if !found {
			break
		}
	}
	return result
}
