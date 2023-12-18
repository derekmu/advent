package day17

import (
	"advent/util"
	"container/heap"
	_ "embed"
	"math"
	"time"
)

type direction byte

const (
	vert direction = 0
	hori           = 1
)

type edge struct {
	heatLoss int
	node     *node
}

type node struct {
	row      int
	col      int
	val      int
	dir      direction
	out      []edge
	open     bool
	closed   bool
	heatLoss int
	index    int
	previous *node
}

type pqueue []*node

func (q pqueue) Len() int {
	return len(q)
}

func (q pqueue) Less(i, j int) bool {
	return q[i].heatLoss < q[j].heatLoss
}

func (q pqueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].index = i
	q[j].index = j
}

func (q *pqueue) Push(x interface{}) {
	n := len(*q)
	no := x.(*node)
	no.index = n
	*q = append(*q, no)
}

func (q *pqueue) Pop() interface{} {
	old := *q
	n := len(old)
	no := old[n-1]
	no.index = -1
	*q = old[0 : n-1]
	return no
}

//go:embed input.txt
var Input []byte

func Run(input []byte) (*util.Result, error) {
	start := time.Now()

	lines := util.ParseInputLines(input)
	nodes := make([][][2]*node, len(lines))
	for row, line := range lines {
		nodes[row] = make([][2]*node, len(line))
		for col, ch := range line {
			nodes[row][col] = [2]*node{
				{
					row:      row,
					col:      col,
					val:      int(ch - '0'),
					dir:      vert,
					out:      make([]edge, 0, 6),
					open:     false,
					closed:   false,
					heatLoss: 0,
					index:    -1,
					previous: nil,
				},
				{
					row:      row,
					col:      col,
					val:      int(ch - '0'),
					dir:      hori,
					out:      make([]edge, 0, 6),
					open:     false,
					closed:   false,
					heatLoss: 0,
					index:    -1,
					previous: nil,
				},
			}
		}
	}
	for row, lineNodes := range nodes {
		for col, n := range lineNodes {
			heatLoss := 0
			nv := n[0]
			nh := n[1]
			for r := row - 1; r >= 0 && r > r-4; r-- {
				heatLoss += nodes[r][col][0].val
				nv.out = append(nv.out, edge{
					heatLoss: heatLoss,
					node:     nodes[r][col][1],
				})
			}
			heatLoss = 0
			for r := row + 1; r < len(nodes) && r < row+4; r++ {
				heatLoss += nodes[r][col][0].val
				nv.out = append(nv.out, edge{
					heatLoss: heatLoss,
					node:     nodes[r][col][1],
				})
			}
			heatLoss = 0
			for c := col - 1; c >= 0 && c > col-4; c-- {
				heatLoss += nodes[row][c][0].val
				nh.out = append(nh.out, edge{
					heatLoss: heatLoss,
					node:     nodes[row][c][0],
				})
			}
			heatLoss = 0
			for c := col + 1; c < len(lineNodes) && c < col+4; c++ {
				heatLoss += nodes[row][c][0].val
				nh.out = append(nh.out, edge{
					heatLoss: heatLoss,
					node:     nodes[row][c][0],
				})
			}
		}
	}

	parse := time.Now()

	part1 := math.MaxInt
	openSet := make(pqueue, 0, len(lines)*len(lines[0]))
	heap.Init(&openSet)
	heap.Push(&openSet, nodes[0][0][0])
	heap.Push(&openSet, nodes[0][0][1])
	for openSet.Len() > 0 {
		n := heap.Pop(&openSet).(*node)
		n.open = false
		n.closed = true
		if n.row == len(lines)-1 && n.col == len(lines[0])-1 {
			part1 = n.heatLoss
			break
		}
		// add new open paths branching from this one
		for _, e := range n.out {
			heatLoss := n.heatLoss + e.heatLoss
			if heatLoss < e.node.heatLoss {
				if e.node.open {
					heap.Remove(&openSet, e.node.index)
				}
				e.node.open = false
				e.node.closed = false
			}
			if !e.node.open && !e.node.closed {
				e.node.heatLoss = heatLoss
				e.node.previous = n
				e.node.open = true
				heap.Push(&openSet, e.node)
			}
		}
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
