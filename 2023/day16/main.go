package day16

import (
	"advent/util"
	_ "embed"
	"runtime"
	"sync"
	"time"
)

type direction byte

const (
	up direction = 1 << iota
	right
	down
	left
)

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

	beams := make(chan beam, len(lines)*2+len(lines[0])*2)
	for r := range lines {
		beams <- beam{
			row: r,
			col: -1,
			dir: right,
		}
		beams <- beam{
			row: r,
			col: len(lines[0]),
			dir: left,
		}
	}
	for c := range lines[0] {
		beams <- beam{
			row: -1,
			col: c,
			dir: down,
		}
		beams <- beam{
			row: len(lines),
			col: c,
			dir: up,
		}
	}
	close(beams)
	part1 := -1
	part2 := -1
	var mux sync.Mutex
	wg := sync.WaitGroup{}
	routinesCount := runtime.NumCPU() * 3
	wg.Add(routinesCount)
	for i := 0; i < routinesCount; i++ {
		go func() {
			p1 := -1
			innerMax := 0
			for b := range beams {
				e := energized(lines, b)
				innerMax = max(innerMax, e)
				if b == (beam{row: 0, col: -1, dir: right}) {
					p1 = e
				}
			}
			mux.Lock()
			if p1 != -1 {
				part1 = p1
			}
			part2 = max(part2, innerMax)
			mux.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()

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
	tileMap := make([][]direction, len(lines))
	for i, line := range lines {
		tileMap[i] = make([]direction, len(line))
	}
	var beams = []beam{b}
	count := 0
	for len(beams) > 0 {
		b = beams[len(beams)-1]
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
			mask := tileMap[b.row][b.col]
			if mask&b.dir != 0 {
				// we've already travelled this tile in this direction
				break
			} else {
				if mask == 0 {
					count++
				}
				tileMap[b.row][b.col] = mask | b.dir
			}
		}
	}
	return count
}
