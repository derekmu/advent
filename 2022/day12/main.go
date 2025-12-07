package day12

import (
	"advent/util"
	"bytes"
	_ "embed"
	"time"
)

//go:embed input.txt
var Input []byte
var Problem = util.Problem{Year: "2022", Day: "12", Runner: Run, Input: Input}

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

func Run(input []byte) (util.Result, error) {
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

	target := bfs1(lines, sp, ep)
	part1 := 0
	for target.from != nil {
		part1++
		target = target.from
	}

	target = bfs2(lines, ep)
	part2 := 0
	for target.from != nil {
		part2++
		target = target.from
	}

	end := time.Now()

	return util.Result{
		Part1:     part1,
		Part2:     part2,
		StartTime: start,
		ParseTime: parse,
		EndTime:   end,
	}, nil
}

func bfs(lines [][]byte, ep point, openNodes []*pathNode) *pathNode {
	rows := len(lines)
	cols := len(lines[0])
	opened := make(map[point]bool, rows*cols)
	for _, n := range openNodes {
		opened[n.point] = true
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

func bfs1(lines [][]byte, sp, ep point) *pathNode {
	rows := len(lines)
	cols := len(lines[0])
	origin := &pathNode{point: sp}
	openNodes := make([]*pathNode, 0, rows*cols)
	openNodes = append(openNodes, origin)
	return bfs(lines, ep, openNodes)
}

func bfs2(lines [][]byte, ep point) *pathNode {
	rows := len(lines)
	cols := len(lines[0])
	openNodes := make([]*pathNode, 0, rows*cols)
	for row, line := range lines {
		for col, ch := range line {
			if ch == 'a' {
				p := point{row: row, col: col}
				pn := &pathNode{point: p}
				openNodes = append(openNodes, pn)
			}
		}
	}
	return bfs(lines, ep, openNodes)
}
