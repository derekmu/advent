package util

import (
	"bytes"
)

var newline = []byte{'\n'}

func ParseInputLines(input []byte) [][]byte {
	count := bytes.Count(input, newline)
	if input[len(input)-1] != '\n' {
		count++
	}
	result := make([][]byte, 0, count)
	var line []byte
	var found bool
	for row := 0; ; row++ {
		line, input, found = bytes.Cut(input, newline)
		if found || len(line) > 0 {
			result = append(result, line)
		}
		if !found {
			break
		}
	}
	return result
}
