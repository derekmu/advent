package day25

import (
	"advent/util"
	_ "embed"
	"math"
	"strings"
	"time"
)

//go:embed input.txt
var Input []byte
var Problem = util.Problem{Year: "2023", Day: "25", Runner: Run, Input: Input}

func parseInput(input []byte) (adjacency [][]int) {
	lines := util.ParseInputLines(input)
	conns := make([][2]int, 0, len(lines)*5)
	names := make(map[string]int, len(lines))
	getIndex := func(name string) int {
		if i, ok := names[name]; ok {
			return i
		} else {
			i = len(names)
			names[name] = i
			return i
		}
	}
	var o string
	for _, lineBytes := range lines {
		line := string(lineBytes)
		o, line, _ = strings.Cut(line, ": ")
		for _, o2 := range strings.Split(line, " ") {
			conns = append(conns, [2]int{getIndex(o), getIndex(o2)})
		}
	}
	adjacency = make([][]int, len(names))
	for i := range adjacency {
		adjacency[i] = make([]int, len(names))
	}
	for _, conn := range conns {
		adjacency[conn[0]][conn[1]] = 1
		adjacency[conn[1]][conn[0]] = 1
	}
	return adjacency
}

func Run(input []byte) (util.Result, error) {
	start := time.Now()

	adjacency := parseInput(input)

	parse := time.Now()

	nodes := globalMinCut(adjacency)

	part1 := len(nodes) * (len(adjacency) - len(nodes))
	part2 := -1

	end := time.Now()

	return util.Result{
		Part1:     part1,
		Part2:     part2,
		StartTime: start,
		ParseTime: parse,
		EndTime:   end,
	}, nil
}

func globalMinCut(mat [][]int) []int {
	bestCut := math.MaxInt
	var bestNodes []int
	n := len(mat)
	co := make([][]int, n)
	for i := 0; i < n; i++ {
		co[i] = []int{i}
	}
	for ph := 1; ph < n; ph++ {
		w := make([]int, n)
		copy(w, mat[0])
		s, t := 0, 0
		for it := 0; it < n-ph; it++ {
			w[t] = math.MinInt
			s, t = t, maxIndex(w)
			for i := 0; i < n; i++ {
				w[i] += mat[t][i]
			}
		}
		cut := w[t] - mat[t][t]
		if cut < bestCut {
			bestCut = cut
			bestNodes = co[t]
		}
		co[s] = append(co[s], co[t]...)
		for i := 0; i < n; i++ {
			mat[s][i] += mat[t][i]
		}
		for i := 0; i < n; i++ {
			mat[i][s] = mat[s][i]
		}
		mat[0][t] = math.MinInt
	}
	return bestNodes
}

func maxIndex(arr []int) int {
	maxIdx := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[maxIdx] {
			maxIdx = i
		}
	}
	return maxIdx
}
