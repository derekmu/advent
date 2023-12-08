package day07

import (
	"advent/util"
	_ "embed"
	"fmt"
	"time"
)

type node struct {
	value      uint32
	left       uint32
	right      uint32
	step       int
	atEndpoint *visitedNode
}

type visited struct {
	value     uint32
	moveIndex int
}

type visitedNode struct {
	visited
	previous    *visitedNode
	next        *visitedNode
	stepsToNext int
}

const (
	origin      = uint32('A')<<16 | uint32('A')<<8 | uint32('A')
	destination = uint32('Z')<<16 | uint32('Z')<<8 | uint32('Z')
)

//go:embed input.txt
var Input []byte

func Run(input []byte) (*util.Result, error) {
	start := time.Now()

	lines := util.ParseInputLines(input)
	moves := lines[0]
	nodeMap := make(map[uint32]*node, len(lines)-2)
	for _, line := range lines[2:] {
		v := uint32(line[0])<<16 | uint32(line[1])<<8 | uint32(line[2])
		l := uint32(line[7])<<16 | uint32(line[8])<<8 | uint32(line[9])
		r := uint32(line[12])<<16 | uint32(line[13])<<8 | uint32(line[14])
		nodeMap[v] = &node{value: v, left: l, right: r}
	}

	parse := time.Now()

	cur := nodeMap[origin]
	part1 := 0
	for cur.value != destination {
		switch moves[part1%len(moves)] {
		case 'L':
			cur = nodeMap[cur.left]
		case 'R':
			cur = nodeMap[cur.right]
		default:
			panic(fmt.Sprintf("unknown move %c", moves[part1%len(moves)]))
		}
		part1++
	}

	starts := make([]*node, 0, 6)
	for _, n := range nodeMap {
		if n.value&0xFF == 'A' {
			starts = append(starts, n)
		}
	}
	for _, ocur := range starts {
		visitedEndpoints := make(map[visited]int, 12)
		step := 0
		var v visited
		var pv *visitedNode
		cur = ocur
		for {
			if cur.value&0xFF == 'Z' {
				v = visited{value: cur.value, moveIndex: step % len(moves)}
				nv := &visitedNode{visited: v, previous: pv}
				if pv != nil {
					pv.next = nv
				}
				pv = nv
				if _, ok := visitedEndpoints[v]; ok {
					break
				}
				visitedEndpoints[v] = step
			}
			switch moves[step%len(moves)] {
			case 'L':
				cur = nodeMap[cur.left]
			case 'R':
				cur = nodeMap[cur.right]
			default:
				panic(fmt.Sprintf("unknown move %c", moves[step%len(moves)]))
			}
			step++
		}
		ocur.step = visitedEndpoints[v] // start at the first endpoint in the loop
		fv := pv.previous
		for fv.visited != pv.visited {
			fv = fv.previous
		}
		fv.previous = pv.previous
		pv.previous.next = fv
		fv.stepsToNext = step - visitedEndpoints[fv.previous.visited]
		fv = fv.next
		for fv.stepsToNext == 0 {
			fv.stepsToNext = visitedEndpoints[fv.next.visited] - visitedEndpoints[fv.visited]
			fv = fv.next
		}
		ocur.atEndpoint = fv
	}
	part2 := 0
	done := false
	for !done {
		var lowestCur *node
		for _, cur := range starts {
			if lowestCur == nil || cur.step < lowestCur.step {
				lowestCur = cur
			}
		}
		lowestCur.step += lowestCur.atEndpoint.stepsToNext
		lowestCur.atEndpoint = lowestCur.atEndpoint.next
		part2 = lowestCur.step
		done = true
		for _, cur := range starts {
			if cur.step != part2 {
				done = false
				break
			}
		}
	}

	end := time.Now()

	return &util.Result{
		Part1:     part1,
		Part2:     part2,
		StartTime: start,
		ParseTime: parse,
		EndTime:   end,
	}, nil
}
