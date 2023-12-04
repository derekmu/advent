package util

import (
	"log"
	"time"
)

func PrintResults(part1, part2 int, start, parse, end time.Time) {
	log.Printf("Answers (part 1 / part 2):        %d / %d", part1, part2)
	log.Printf("Timing (parse + execute = total): %v + %v = %v", parse.Sub(start), end.Sub(parse), end.Sub(start))
}
