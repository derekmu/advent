package day23

import (
	"advent/util"
	_ "embed"
	"time"
)

//go:embed input.txt
var Input []byte
var Problem = util.Problem{Year: "2023", Day: "XX", Runner: Run, Input: Input}

type path struct {
	stepCount int
	row       int
	col       int
	visited   [][]bool
}

type direnum byte

const (
	up direnum = iota
	right
	down
	left
)

type path2 struct {
	e       *edge
	row     int
	col     int
	exclude direnum
}

type direction struct {
	drow    int
	dcol    int
	slope   byte
	dir     direnum
	exclude direnum
}

var dirs = []direction{
	{dcol: 1, slope: '>', dir: right, exclude: left},
	{drow: 1, slope: 'v', dir: down, exclude: up},
	{dcol: -1, slope: '<', dir: left, exclude: right},
	{drow: -1, slope: '^', dir: up, exclude: down},
}

type node struct {
	row     int
	col     int
	edges   []*edge
	visited bool
}

type edge struct {
	stepCount int
	node1     *node
	node2     *node
}

func Run(input []byte) (util.Result, error) {
	start := time.Now()

	lines := util.ParseInputLines(input)

	parse := time.Now()

	part1 := solvePart1(lines)
	part2 := solvePart2(lines)

	end := time.Now()

	return util.Result{
		Part1:     part1,
		Part2:     part2,
		StartTime: start,
		ParseTime: parse,
		EndTime:   end,
	}, nil
}

func solvePart1(lines [][]byte) int {
	width := len(lines)
	height := len(lines[0])
	empty := make([][]bool, height)
	for i := range empty {
		empty[i] = make([]bool, width)
	}
	part1 := -1
	paths := make([]path, 0, 50)
	paths = append(paths, path{stepCount: 0, row: 0, col: 1, visited: copyVisited(empty)})
	for len(paths) > 0 {
		p := paths[len(paths)-1]
		paths = paths[:len(paths)-1]
		moved := true
		for moved {
			moved = false
			p.visited[p.row][p.col] = true
			if p.row == height-1 && p.col == width-2 {
				part1 = max(part1, p.stepCount)
				break
			}
			switch lines[p.row][p.col] {
			case '>':
				p.col++
				p.stepCount++
				moved = true
				continue
			case '<':
				p.col--
				p.stepCount++
				moved = true
				continue
			case '^':
				p.row--
				p.stepCount++
				moved = true
				continue
			case 'v':
				p.row++
				p.stepCount++
				moved = true
				continue
			}
			op := p
			for _, d := range dirs {
				row := op.row + d.drow
				col := op.col + d.dcol
				if row >= 0 && row < height && col >= 0 && col < width && (lines[row][col] == '.' || lines[row][col] == d.slope) && !p.visited[row][col] {
					if moved {
						paths = append(paths, path{stepCount: op.stepCount + 1, row: row, col: col, visited: copyVisited(p.visited)})
					} else {
						p.row = row
						p.col = col
						p.stepCount++
						moved = true
					}
				}
			}
		}
	}
	return part1
}

func copyVisited(visited [][]bool) [][]bool {
	n := make([][]bool, len(visited))
	for row, visits := range visited {
		n[row] = make([]bool, len(visits))
		copy(n[row], visits)
	}
	return n
}

func solvePart2(lines [][]byte) int {
	width := len(lines)
	height := len(lines[0])
	nodeMap := make([][]*node, height)
	for row := range nodeMap {
		nodeMap[row] = make([]*node, width)
	}
	// start
	nodeMap[0][1] = &node{
		row:   0,
		col:   1,
		edges: []*edge{{stepCount: 1}},
	}
	nodeMap[0][1].edges[0].node1 = nodeMap[0][1]
	// end
	nodeMap[height-1][width-2] = &node{
		row:   height - 1,
		col:   width - 2,
		edges: make([]*edge, 0, 4),
	}

	paths := make([]path2, 0, 50)
	paths = append(paths, path2{e: nodeMap[0][1].edges[0], row: 1, col: 1, exclude: up})

	for len(paths) > 0 {
		p := paths[len(paths)-1]
		paths = paths[:len(paths)-1]
		branchCount := 1
		for branchCount == 1 {
			branchCount = 0
			for _, d := range dirs {
				if d.dir == p.exclude {
					continue
				}
				row := p.row + d.drow
				col := p.col + d.dcol
				if row >= 0 && row < height && col >= 0 && col < width && lines[row][col] != '#' {
					branchCount++
				}
			}
			if branchCount == 0 {
				// should be the finish
				enode := nodeMap[p.row][p.col]
				p.e.node2 = enode
				enode.edges = append(enode.edges, p.e)
			} else if branchCount == 1 {
				for _, d := range dirs {
					if d.dir == p.exclude {
						continue
					}
					row := p.row + d.drow
					col := p.col + d.dcol
					if row >= 0 && row < height && col >= 0 && col < width && lines[row][col] != '#' {
						p.row = row
						p.col = col
						p.e.stepCount++
						p.exclude = d.exclude
						break
					}
				}
			} else if branchCount > 1 {
				enode := nodeMap[p.row][p.col]
				if enode == nil {
					enode = &node{
						row:     p.row,
						col:     p.col,
						edges:   make([]*edge, 0, branchCount),
						visited: false,
					}
					nodeMap[p.row][p.col] = enode
					enode.edges = append(enode.edges, p.e)
					for _, d := range dirs {
						if d.dir == p.exclude {
							continue
						}
						row := p.row + d.drow
						col := p.col + d.dcol
						if row >= 0 && row < height && col >= 0 && col < width && lines[row][col] != '#' {
							e := &edge{
								stepCount: 1,
								node1:     enode,
							}
							enode.edges = append(enode.edges, e)
							paths = append(paths, path2{e: e, row: row, col: col, exclude: d.exclude})
						}
					}
				}
				p.e.node2 = enode
			}
		}
	}
	if result, ok := longestPath(nodeMap[0][1], nodeMap[height-1][width-2], 0); ok {
		return result
	} else {
		panic("did not find end point")
	}
}

func longestPath(n, target *node, stepCount int) (result int, ok bool) {
	if n == target {
		return stepCount, true
	}
	n.visited = true
	for _, e := range n.edges {
		if !e.node1.visited {
			if steps, found := longestPath(e.node1, target, stepCount+e.stepCount); found {
				result = max(result, steps)
				ok = true
			}
		} else if !e.node2.visited {
			if steps, found := longestPath(e.node2, target, stepCount+e.stepCount); found {
				result = max(result, steps)
				ok = true
			}
		}
	}
	n.visited = false
	return result, ok
}
