package util

import (
	"bytes"
)

var newline = []byte{'\n'}

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
