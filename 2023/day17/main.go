package day17

import (
	"advent/util"
	"container/heap"
	_ "embed"
	"time"
)

var Problem = util.Problem{Year: "2023", Day: "17", Runner: Run, Input: Input}

type edge struct {
	heatLoss int
	node     *node
}

type node struct {
	row      int
	col      int
	val      int
	out1     []edge
	out2     []edge
	open     bool
	closed   bool
	heatLoss int
	index    int
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
					row:  row,
					col:  col,
					val:  int(ch - '0'),
					out1: make([]edge, 0, 6),
					out2: make([]edge, 0, 14),
				},
				{
					row:  row,
					col:  col,
					val:  int(ch - '0'),
					out1: make([]edge, 0, 6),
					out2: make([]edge, 0, 14),
				},
			}
		}
	}
	for row, lineNodes := range nodes {
		for col, n := range lineNodes {
			heatLoss := 0
			nv := n[0]
			nh := n[1]
			for r := row - 1; r >= 0 && r > r-11; r-- {
				heatLoss += nodes[r][col][0].val
				if r > row-4 {
					nv.out1 = append(nv.out1, edge{
						heatLoss: heatLoss,
						node:     nodes[r][col][1],
					})
				} else {
					nv.out2 = append(nv.out2, edge{
						heatLoss: heatLoss,
						node:     nodes[r][col][1],
					})
				}
			}
			heatLoss = 0
			for r := row + 1; r < len(nodes) && r < row+11; r++ {
				heatLoss += nodes[r][col][0].val
				if r < row+4 {
					nv.out1 = append(nv.out1, edge{
						heatLoss: heatLoss,
						node:     nodes[r][col][1],
					})
				} else {
					nv.out2 = append(nv.out2, edge{
						heatLoss: heatLoss,
						node:     nodes[r][col][1],
					})
				}
			}
			heatLoss = 0
			for c := col - 1; c >= 0 && c > col-11; c-- {
				heatLoss += nodes[row][c][0].val
				if c > col-4 {
					nh.out1 = append(nh.out1, edge{
						heatLoss: heatLoss,
						node:     nodes[row][c][0],
					})
				} else {
					nh.out2 = append(nh.out2, edge{
						heatLoss: heatLoss,
						node:     nodes[row][c][0],
					})
				}
			}
			heatLoss = 0
			for c := col + 1; c < len(lineNodes) && c < col+11; c++ {
				heatLoss += nodes[row][c][0].val
				if c < col+4 {
					nh.out1 = append(nh.out1, edge{
						heatLoss: heatLoss,
						node:     nodes[row][c][0],
					})
				} else {
					nh.out2 = append(nh.out2, edge{
						heatLoss: heatLoss,
						node:     nodes[row][c][0],
					})
				}
			}
		}
	}

	parse := time.Now()

	part1 := solve(nodes, false)
	part2 := solve(nodes, true)

	end := time.Now()

	return &util.Result{
		Part1:     part1,
		Part2:     part2,
		StartTime: start,
		ParseTime: parse,
		EndTime:   end,
	}, nil
}

func solve(nodes [][][2]*node, part2 bool) int {
	for row := range nodes {
		for col := range nodes[row] {
			for _, n := range nodes[row][col] {
				n.open = false
				n.closed = false
				n.heatLoss = 0
				n.index = -1
			}
		}
	}
	openSet := make(pqueue, 0, len(nodes)*len(nodes[0]))
	heap.Init(&openSet)
	heap.Push(&openSet, nodes[0][0][0])
	heap.Push(&openSet, nodes[0][0][1])
	for openSet.Len() > 0 {
		n := heap.Pop(&openSet).(*node)
		n.open = false
		n.closed = true
		if n.row == len(nodes)-1 && n.col == len(nodes[0])-1 {
			return n.heatLoss
		}
		// add new open paths branching from this one
		outs := n.out1
		if part2 {
			outs = n.out2
		}
		for _, e := range outs {
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
				e.node.open = true
				heap.Push(&openSet, e.node)
			}
		}
	}
	panic("solution not found")
}
