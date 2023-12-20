package day14

import (
	"advent/util"
	"bytes"
	_ "embed"
	"time"
)

//go:embed input.txt
var Input []byte
var Problem = util.Problem{Year: "2022", Day: "14", Runner: Run, Input: Input}

type point struct {
	col int
	row int
}

type rocks struct {
	paths                  [][]point
	minCol, maxCol, maxRow int
}

func parseInput(input []byte) (rs rocks) {
	lines := util.ParseInputLines(input)
	rs.paths = make([][]point, 0, len(lines))
	rs.minCol = 500
	rs.maxCol = 500
	rs.maxRow = 0
	for _, line := range lines {
		pointStrings := bytes.Split(line, []byte(" -> "))
		rockPath := make([]point, 0, len(pointStrings))
		for _, pointString := range pointStrings {
			colString, rowString, _ := bytes.Cut(pointString, []byte(","))
			p := point{
				col: util.Btoi(colString),
				row: util.Btoi(rowString),
			}
			rockPath = append(rockPath, p)
			rs.minCol = min(rs.minCol, p.col)
			rs.maxCol = max(rs.maxCol, p.col)
			rs.maxRow = max(rs.maxRow, p.row)
		}
		rs.paths = append(rs.paths, rockPath)
	}
	return rs
}

func Run(input []byte) (*util.Result, error) {
	start := time.Now()

	rs := parseInput(input)

	parse := time.Now()

	m := makeMap(rs)

	part1 := 0
	for {
		sand := point{500 - rs.minCol + 1, 0}
		for sand.row < rs.maxRow {
			if m[sand.row+1][sand.col] == ' ' {
				sand.row++
			} else if m[sand.row+1][sand.col-1] == ' ' {
				sand.col--
				sand.row++
			} else if m[sand.row+1][sand.col+1] == ' ' {
				sand.col++
				sand.row++
			} else {
				break
			}
		}
		if sand.row >= rs.maxRow {
			break
		} else {
			m[sand.row][sand.col] = 'O'
			part1++
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

func makeMap(rs rocks) [][]byte {
	m := make([][]byte, rs.maxRow+1)
	for row := range m {
		m[row] = make([]byte, rs.maxCol-rs.minCol+3)
		for i := range m[row] {
			m[row][i] = ' '
		}
	}
	for _, path := range rs.paths {
		for i := 0; i < len(path)-1; i++ {
			p1 := path[i]
			p1.col = p1.col - rs.minCol + 1
			p2 := path[i+1]
			p2.col = p2.col - rs.minCol + 1
			m[p1.row][p1.col] = '#'
			for p1.row < p2.row {
				p1.row++
				m[p1.row][p1.col] = '#'
			}
			for p1.row > p2.row {
				p1.row--
				m[p1.row][p1.col] = '#'
			}
			for p1.col < p2.col {
				p1.col++
				m[p1.row][p1.col] = '#'
			}
			for p1.col > p2.col {
				p1.col--
				m[p1.row][p1.col] = '#'
			}
		}
	}
	return m
}
