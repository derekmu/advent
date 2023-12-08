package util

import (
	"fmt"
)

func Btoi(bytes []byte) int {
	var result int
	negate := bytes[0] == '-'
	if negate {
		bytes = bytes[1:]
	}
	for _, b := range bytes {
		if b < '0' || b > '9' {
			panic(fmt.Sprintf("invalid character %d %c", b, b))
		}
		result = result*10 + int(b-'0')
	}
	if negate {
		return -result
	}
	return result
}
