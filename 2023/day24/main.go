package day24

import (
	"advent/util"
	"bytes"
	_ "embed"
	"slices"
	"time"
)

//go:embed input.txt
var Input []byte
var Problem = util.Problem{Year: "2023", Day: "24", Runner: Run, Input: Input}

type stone struct {
	x, y, z    float64
	dx, dy, dz float64
}

type vector3 struct {
	x, y, z float64
}

func parseInput(input []byte) (stones []stone, minint, maxint float64) {
	lines := util.ParseInputLines(input)
	line1 := lines[0]
	minints, maxints, _ := bytes.Cut(line1, []byte(", "))
	minint = float64(util.Btoi(minints))
	maxint = float64(util.Btoi(maxints))
	lines = lines[1:]
	stones = make([]stone, len(lines))
	var x, y, z, dx, dy, dz []byte
	for i, line := range lines {
		x, line, _ = bytes.Cut(line, []byte(", "))
		y, line, _ = bytes.Cut(line, []byte(", "))
		z, line, _ = bytes.Cut(line, []byte(" @ "))
		dx, line, _ = bytes.Cut(line, []byte(", "))
		dy, dz, _ = bytes.Cut(line, []byte(", "))
		s := stone{
			x:  float64(util.Btoi(x)),
			y:  float64(util.Btoi(y)),
			z:  float64(util.Btoi(z)),
			dx: float64(util.Btoi(dx)),
			dy: float64(util.Btoi(dy)),
			dz: float64(util.Btoi(dz)),
		}
		stones[i] = s
	}
	return stones, minint, maxint
}

func Run(input []byte) (*util.Result, error) {
	start := time.Now()

	stones, mini, maxi := parseInput(input)

	parse := time.Now()

	part1 := 0
	for i, s1 := range stones {
		for j := i + 1; j < len(stones); j++ {
			s2 := stones[j]
			det := s1.dx*s2.dy - s1.dy*s2.dx
			if det == 0 {
				continue
			}

			t := ((s2.x-s1.x)*s2.dy - (s2.y-s1.y)*s2.dx) / det
			u := ((s2.x-s1.x)*s1.dy - (s2.y-s1.y)*s1.dx) / det

			if t >= 0 && u >= 0 {
				xi := s1.x + t*s1.dx
				yi := s1.y + t*s1.dy
				if xi >= mini && xi <= maxi && yi >= mini && yi <= maxi {
					part1++
				}
			}
		}
	}

	var maybeX, maybeY, maybeZ []int
	for i := 0; i < len(stones)-1; i++ {
		for j := i + 1; j < len(stones); j++ {
			a, b := stones[i], stones[j]
			if a.dx == b.dx {
				nextMaybe := findMatchingVel(int(b.x-a.x), int(a.dx))
				if len(maybeX) == 0 {
					maybeX = nextMaybe
				} else {
					maybeX = getIntersect(maybeX, nextMaybe)
				}
			}
			if a.dy == b.dy {
				nextMaybe := findMatchingVel(int(b.y-a.y), int(a.dy))
				if len(maybeY) == 0 {
					maybeY = nextMaybe
				} else {
					maybeY = getIntersect(maybeY, nextMaybe)
				}
			}
			if a.dz == b.dz {
				nextMaybe := findMatchingVel(int(b.z-a.z), int(a.dz))
				if len(maybeZ) == 0 {
					maybeZ = nextMaybe
				} else {
					maybeZ = getIntersect(maybeZ, nextMaybe)
				}
			}
		}
	}

	part2 := -1
	if len(maybeX) == len(maybeY) && len(maybeY) == len(maybeZ) && len(maybeZ) == 1 {
		// only one possible velocity in all dimensions
		rockVel := vector3{float64(maybeX[0]), float64(maybeY[0]), float64(maybeZ[0])}
		hailStoneA, hailStoneB := stones[0], stones[1]
		mA := (hailStoneA.dy - rockVel.y) / (hailStoneA.dx - rockVel.x)
		mB := (hailStoneB.dy - rockVel.y) / (hailStoneB.dx - rockVel.x)
		cA := hailStoneA.y - (mA * hailStoneA.x)
		cB := hailStoneB.y - (mB * hailStoneB.x)
		xPos := (cB - cA) / (mA - mB)
		yPos := mA*xPos + cA
		t := (xPos - hailStoneA.x) / (hailStoneA.dx - rockVel.x)
		zPos := hailStoneA.z + (hailStoneA.dz-rockVel.z)*t
		part2 = int(xPos + yPos + zPos)
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

func findMatchingVel(dvel, pv int) []int {
	var match []int
	for v := -1000; v < 1000; v++ {
		if v != pv && dvel%(v-pv) == 0 {
			match = append(match, v)
		}
	}
	return match
}

func getIntersect(a, b []int) []int {
	var result []int
	for _, val := range a {
		if slices.Contains(b, val) {
			result = append(result, val)
		}
	}
	return result
}
