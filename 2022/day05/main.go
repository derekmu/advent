package day05

import (
	"advent/util"
	"bytes"
	"log"
)

func Run(input []byte) error {
	lines := util.ParseInputLines(input)
	lineI := 0

	// read stacks
	var stacks1 [][]uint8
	var lastLine []byte
	for ; lineI < len(lines); lineI++ {
		line := lines[lineI]
		if len(line) == 0 {
			break
		}
		for stackIndex := 0; stackIndex < (len(lastLine)+1)/4; stackIndex++ {
			if lastLine[stackIndex*4] == '[' {
				for stackIndex >= len(stacks1) {
					stacks1 = append(stacks1, nil)
				}
				stacks1[stackIndex] = append(stacks1[stackIndex], lastLine[stackIndex*4+1])
			}
		}
		lastLine = line
	}

	// invert and copy for part 2
	stacks2 := make([][]uint8, len(stacks1))
	for si, stack := range stacks1 {
		for i := 0; i < len(stack)/2; i++ {
			stack[i], stack[len(stack)-1-i] = stack[len(stack)-1-i], stack[i]
		}
		stacks2[si] = make([]uint8, len(stack))
		copy(stacks2[si], stack)
	}

	lineI++

	// do moves
	for ; lineI < len(lines); lineI++ {
		line := lines[lineI]
		parts := bytes.Split(line, []byte(" "))
		moves := util.Btoi(parts[1])
		fromIndex := util.Btoi(parts[3])
		fromIndex--
		toIndex := util.Btoi(parts[5])
		toIndex--
		for n := 0; n < moves; n++ {
			stacks1[toIndex] = append(stacks1[toIndex], stacks1[fromIndex][len(stacks1[fromIndex])-1-n])
			stacks2[toIndex] = append(stacks2[toIndex], stacks2[fromIndex][len(stacks2[fromIndex])-moves+n])
		}
		stacks1[fromIndex] = stacks1[fromIndex][:len(stacks1[fromIndex])-moves]
		stacks2[fromIndex] = stacks2[fromIndex][:len(stacks2[fromIndex])-moves]
	}

	// assemble answer
	result1 := ""
	result2 := ""
	for i := 0; i < len(stacks1); i++ {
		result1 = result1 + string(stacks1[i][len(stacks1[i])-1])
		result2 = result2 + string(stacks2[i][len(stacks2[i])-1])
	}

	log.Printf("The crates at the top of each stack for part 1 are %s", result1)
	log.Printf("The crates at the top of each stack for part 2 are %s", result2)

	return nil
}
