package day16

import (
	"advent/util"
	_ "embed"
	"time"
)

type direction byte

const (
	up direction = 1 << iota
	right
	down
	left
)

type tile struct {
	row int
	col int
}

type beam struct {
	row int
	col int
	dir direction
}

//go:embed input.txt
var Input []byte

func Run(input []byte) (*util.Result, error) {
	start := time.Now()

	lines := util.ParseInputLines(input)

	parse := time.Now()

	part1 := energized(lines, beam{
		row: 0,
		col: -1,
		dir: right,
	})
	part2 := -1
	for r := range lines {
		e := energized(lines, beam{
			row: r,
			col: -1,
			dir: right,
		})
		part2 = max(part2, e)
		e = energized(lines, beam{
			row: r,
			col: len(lines[0]),
			dir: left,
		})
		part2 = max(part2, e)
	}
	for c := range lines[0] {
		e := energized(lines, beam{
			row: -1,
			col: c,
			dir: down,
		})
		part2 = max(part2, e)
		e = energized(lines, beam{
			row: len(lines),
			col: c,
			dir: up,
		})
		part2 = max(part2, e)
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

func energized(lines [][]byte, b beam) int {
	tileMap := make(map[tile]direction, len(lines)*len(lines[0]))
	var beams = []beam{b}
	for len(beams) > 0 {
		b := beams[len(beams)-1]
		beams = beams[:len(beams)-1]
		for {
			switch b.dir {
			case up:
				b.row = b.row - 1
			case right:
				b.col = b.col + 1
			case down:
				b.row = b.row + 1
			case left:
				b.col = b.col - 1
			default:
				panic("unknown direction")
			}
			if b.row < 0 || b.row >= len(lines) || b.col < 0 || b.col >= len(lines[0]) {
				break
			}
			switch lines[b.row][b.col] {
			case '.':
				// beam continues
			case '|':
				switch b.dir {
				case up, down:
				// beam continues
				case right, left:
					b.dir = up
					beams = append(beams, beam{
						row: b.row,
						col: b.col,
						dir: down,
					})
				default:
					panic("unknown direction")
				}
			case '-':
				switch b.dir {
				case right, left:
				// beam continues
				case up, down:
					b.dir = left
					beams = append(beams, beam{
						row: b.row,
						col: b.col,
						dir: right,
					})
				default:
					panic("unknown direction")
				}
			case '\\':
				switch b.dir {
				case right:
					b.dir = down
				case left:
					b.dir = up
				case up:
					b.dir = left
				case down:
					b.dir = right
				default:
					panic("unknown direction")
				}
			case '/':
				switch b.dir {
				case right:
					b.dir = up
				case left:
					b.dir = down
				case up:
					b.dir = right
				case down:
					b.dir = left
				default:
					panic("unknown direction")
				}
			default:
				panic("unknown input character")
			}
			if mask, ok := tileMap[tile{row: b.row, col: b.col}]; ok {
				if mask&b.dir != 0 {
					// we've already travelled this tile in this direction
					break
				} else {
					tileMap[tile{row: b.row, col: b.col}] = mask | b.dir
				}
			} else {
				tileMap[tile{row: b.row, col: b.col}] = b.dir
			}
		}
	}
	return len(tileMap)
}
