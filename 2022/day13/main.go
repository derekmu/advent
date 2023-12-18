package day13

import (
	"advent/util"
	_ "embed"
	"slices"
	"time"
)

var Problem = util.Problem{Year: "2022", Day: "13", Runner: Run, Input: Input}

type node struct {
	value  byte
	list   []*node
	parent *node
}

//go:embed input.txt
var Input []byte

func Run(input []byte) (*util.Result, error) {
	start := time.Now()

	lines := util.ParseInputLines(input)
	nodePairs := make([][2]*node, 0, (len(lines)+1)/3)
	nodes := make([]*node, 0, (len(lines)+1)/3*2+2)
	divider1 := parseNode([]byte("[[2]]"))
	divider2 := parseNode([]byte("[[6]]"))
	nodes = append(nodes, divider1)
	nodes = append(nodes, divider2)
	for i := 0; i < len(lines); i += 3 {
		node1 := parseNode(lines[i])
		node2 := parseNode(lines[i+1])
		nodePairs = append(nodePairs, [2]*node{node1, node2})
		nodes = append(nodes, node1, node2)
	}

	parse := time.Now()

	part1 := 0
	for i, np := range nodePairs {
		if compare(np[0], np[1]) < 0 {
			part1 += i + 1
		}
	}
	slices.SortFunc(nodes, compare)
	i1 := slices.Index(nodes, divider1)
	i2 := slices.Index(nodes, divider2)
	part2 := (i1 + 1) * (i2 + 1)

	end := time.Now()

	return &util.Result{
		Part1:     part1,
		Part2:     part2,
		StartTime: start,
		ParseTime: parse,
		EndTime:   end,
	}, nil
}

func compare(ln *node, rn *node) int {
	if ln.value != 0 && rn.value != 0 {
		// both integers
		if ln.value < rn.value {
			return -1 // good
		} else if ln.value > rn.value {
			return 1 // bad
		} else {
			return 0 // same
		}
	} else if ln.value == 0 && rn.value == 0 {
		// both lists
		for i := 0; ; i++ {
			if i >= len(ln.list) && i >= len(rn.list) {
				return 0
			} else if i < len(ln.list) && i < len(rn.list) {
				c := compare(ln.list[i], rn.list[i])
				if c != 0 {
					return c
				}
			} else if i < len(ln.list) {
				return 1 // bad
			} else {
				return -1 // good
			}
		}
	} else if ln.value != 0 {
		// left is an integer, right is a list
		return compare(&node{list: []*node{ln}}, rn)
	} else {
		// left is a list, right is an integer
		return compare(ln, &node{list: []*node{rn}})
	}
}

func parseNode(bytes []byte) *node {
	root := &node{}
	current := root
	for i := 1; i < len(bytes)-1; i++ {
		c := bytes[i]
		switch {
		case c == '[':
			neww := &node{parent: current}
			current.list = append(current.list, neww)
			current = neww
		case c == ']':
			current = current.parent
		case c >= '0' && c <= ':':
			current.list = append(current.list, &node{value: c})
		default:
		}
	}
	return root
}
