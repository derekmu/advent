package main

import (
	"advent/util"
	"bytes"
	"time"
)

type point struct {
	row int
	col int
}

var dirs = []point{
	{row: 1},
	{row: -1},
	{col: 1},
	{col: -1},
}

type pathNode struct {
	point point
	from  *pathNode
}

func main() {
	input := util.ReadInput()

	start := time.Now()

	lines := util.ParseInputLines(input)
	sp := point{row: -1, col: -1}
	ep := point{row: -1, col: -1}
	for row, line := range lines {
		if sp.col < 0 {
			sp.col = bytes.IndexByte(line, 'S')
			if sp.col >= 0 {
				sp.row = row
				line[sp.col] = 'a'
			}
		}
		if ep.col < 0 {
			ep.col = bytes.IndexByte(line, 'E')
			if ep.col >= 0 {
				ep.row = row
				line[ep.col] = 'z'
			}
		}
		if sp.col >= 0 && ep.col >= 0 {
			break
		}
	}

	parse := time.Now()

	target := bfs(lines, sp, ep)
	part1 := 0
	for target.from != nil {
		part1++
		target = target.from
	}

	target = bfsa(lines, ep)
	part2 := 0
	for target.from != nil {
		part2++
		target = target.from
	}

	end := time.Now()

	util.PrintResults(part1, part2, start, parse, end)
}

func bfs(lines [][]byte, sp, ep point) *pathNode {
	rows := len(lines)
	cols := len(lines[0])
	origin := &pathNode{point: sp}
	openNodes := make([]*pathNode, 0, rows*cols)
	opened := make(map[point]bool, rows*cols)
	openNodes = append(openNodes, origin)
	opened[sp] = true
	for len(openNodes) > 0 {
		node := openNodes[0]
		openNodes = openNodes[1:]
		for _, dir := range dirs {
			p := point{node.point.row + dir.row, node.point.col + dir.col}
			if p.row >= 0 && p.row < rows && p.col >= 0 && p.col < cols {
				dv := int(lines[p.row][p.col]) - int(lines[node.point.row][node.point.col])
				if dv <= 1 {
					if _, ok := opened[p]; !ok {
						if p == ep {
							return &pathNode{point: p, from: node}
						} else {
							openNodes = append(openNodes, &pathNode{point: p, from: node})
							opened[p] = true
						}
					}
				}
			}
		}
	}
	return nil
}

func bfsa(lines [][]byte, ep point) *pathNode {
	rows := len(lines)
	cols := len(lines[0])
	openNodes := make([]*pathNode, 0, rows*cols)
	opened := make(map[point]bool, rows*cols)
	for row, line := range lines {
		for col, ch := range line {
			if ch == 'a' {
				p := point{row: row, col: col}
				pn := &pathNode{point: p}
				openNodes = append(openNodes, pn)
				opened[p] = true
			}
		}
	}
	for len(openNodes) > 0 {
		node := openNodes[0]
		openNodes = openNodes[1:]
		for _, dir := range dirs {
			p := point{node.point.row + dir.row, node.point.col + dir.col}
			if p.row >= 0 && p.row < rows && p.col >= 0 && p.col < cols {
				dv := int(lines[p.row][p.col]) - int(lines[node.point.row][node.point.col])
				if dv <= 1 {
					if _, ok := opened[p]; !ok {
						if p == ep {
							return &pathNode{point: p, from: node}
						} else {
							openNodes = append(openNodes, &pathNode{point: p, from: node})
							opened[p] = true
						}
					}
				}
			}
		}
	}
	return nil
}
