package day13

import (
	"advent/util"
	"slices"
	"time"
)

type node struct {
	value  byte
	list   []*node
	parent *node
}

func Run(input []byte) error {
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

	util.PrintResults(part1, part2, start, parse, end)

	return nil
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
	i := 1
	for i < len(bytes)-1 {
		c := bytes[i]
		if c == '[' {
			neww := &node{parent: current}
			current.list = append(current.list, neww)
			current = neww
		} else if c == ']' {
			current = current.parent
		} else if isDigit(c) {
			current.list = append(current.list, &node{value: c})
		} else if c == ',' {
		}
		i++
	}
	return root
}

func isDigit(c byte) bool {
	return c >= '0' && c <= ':'
}
