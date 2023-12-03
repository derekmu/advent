package util

import (
	"log"
)

func Btoi(bytes []byte) int {
	var result int
	for _, b := range bytes {
		if b < '0' || b > '9' {
			log.Panicf("invalid character %d %c", b, b)
		}
		result = result*10 + int(b-'0')
	}
	return result
}
