package day09

import (
	"advent/util"
	"bytes"
	_ "embed"
	"time"
)

//go:embed input.txt
var Input []byte

func parseInput(input []byte) [][]int {
	lines := util.ParseInputLines(input)
	output := make([][]int, 0, len(lines))
	for _, line := range lines {
		nums := bytes.Split(line, []byte(" "))
		out := make([]int, 0, len(nums)+1)
		for _, num := range nums {
			out = append(out, util.Btoi(num))
		}
		output = append(output, out)
	}
	return output
}

func Run(input []byte) (*util.Result, error) {
	start := time.Now()

	sequences := parseInput(input)

	parse := time.Now()

	layers := make([][]int, 19)
	part1 := 0
	var layer, prev []int
	for _, seq := range sequences {
		li := 0
		done := false
		prev = seq
		for !done {
			if li >= len(layers) {
				layer = make([]int, 0, 22)
				layers = append(layers, layer)
			} else {
				layer = layers[li][:0]
			}
			done = true
			for i := 0; i < len(prev)-1; i++ {
				d := prev[i+1] - prev[i]
				layer = append(layer, d)
				done = done && d == 0
			}
			layers[li] = layer
			prev = layer
			li++
		}
		li--
		for ; li >= 0; li-- {
			layer = layers[li]
			if li == 0 {
				seq = append(seq, seq[len(seq)-1]+layer[len(layer)-1])
			} else {
				layers[li-1] = append(layers[li-1], layers[li-1][len(layers[li-1])-1]+layer[len(layer)-1])
			}
		}
		part1 += seq[len(seq)-1]
	}
	part2 := -1

	end := time.Now()

	return &util.Result{
		Part1:     part1,
		Part2:     part2,
		StartTime: start,
		ParseTime: parse,
		EndTime:   end,
	}, nil
}
