package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"time"
)

type num struct {
	row      int
	start    int
	end      int
	adjacent bool
}

type gearIndex struct {
	row int
	col int
}

type gear struct {
	row  int
	col  int
	nums []int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	start := time.Now()

	var lines []string
	var nums []*num
	gearMap := make(map[gearIndex]*gear, 10)
	for row := 0; scanner.Scan(); row++ {
		line := scanner.Text()
		lines = append(lines, line)
		var cnum *num
		for col, ch := range line {
			if ch >= '0' && ch <= '9' {
				if cnum != nil && cnum.end == col-1 {
					cnum.end = col
				} else {
					cnum = &num{
						row:   row,
						start: col,
						end:   col,
					}
					nums = append(nums, cnum)
				}
			} else if ch == '*' {
				gearMap[gearIndex{row: row, col: col}] = &gear{row: row, col: col}
			}
		}
	}
	parse := time.Now()

	sum1 := 0
	for _, n := range nums {
		for row := n.row - 1; row <= n.row+1; row++ {
			if row < 0 || row >= len(lines) {
				continue
			}
			for col := n.start - 1; col <= n.end+1; col++ {
				if row == n.row && col >= n.start && col <= n.end {
					continue
				}
				if col < 0 || col >= len(lines[row]) {
					continue
				}
				ch := lines[row][col]
				switch ch {
				case '.', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				default:
					i, err := strconv.Atoi(lines[n.row][n.start : n.end+1])
					if err != nil {
						log.Panic(err)
					}
					sum1 += i
					if ch == '*' {
						g := gearMap[gearIndex{row: row, col: col}]
						g.nums = append(g.nums, i)
					}
				}
			}
		}
	}

	sum2 := 0
	for _, g := range gearMap {
		if len(g.nums) == 2 {
			sum2 += g.nums[0] * g.nums[1]
		}
	}
	end := time.Now()

	log.Printf("The sum of the part numbers for part 1 is %d", sum1)
	log.Printf("The sum of the part numbers for part 2 is %d", sum2)
	log.Printf("Parse time was %v", parse.Sub(start))
	log.Printf("Execute time was %v", end.Sub(parse))
}
