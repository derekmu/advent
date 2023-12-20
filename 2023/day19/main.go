package day19

import (
	"advent/util"
	"bytes"
	_ "embed"
	"time"
)

//go:embed input.txt
var Input []byte
var Problem = util.Problem{Year: "2023", Day: "19", Runner: Run, Input: Input}

type rule struct {
	attr       byte
	comp       byte
	value      int
	workflowId uint32
}

type workflow struct {
	rules []rule
}

type part struct {
	a int
	m int
	s int
	x int
}

type partRange struct {
	workflowId uint32
	p1, p2     part
}

func (pr *partRange) apply(r rule) partRange {
	pr2 := *pr
	pr2.workflowId = r.workflowId
	switch r.comp {
	case '>':
		pr2.p1.set(r.attr, max(pr2.p1.get(r.attr), r.value+1))
		pr.p2.set(r.attr, min(pr.p2.get(r.attr), r.value))
	case '<':
		pr2.p2.set(r.attr, min(pr2.p2.get(r.attr), r.value-1))
		pr.p1.set(r.attr, max(pr.p1.get(r.attr), r.value))
	case 0:
		// final rule, nothing to do
	default:
		panic("unexpected comparison")
	}
	return pr2
}

func (pr *partRange) combinations() int {
	if pr.p1.a > pr.p2.a || pr.p1.m > pr.p2.m || pr.p1.s > pr.p2.s || pr.p1.x > pr.p2.x {
		return 0
	}
	return (pr.p2.a - pr.p1.a + 1) * (pr.p2.m - pr.p1.m + 1) * (pr.p2.s - pr.p1.s + 1) * (pr.p2.x - pr.p1.x + 1)
}

func (p *part) get(n byte) int {
	switch n {
	case 'a':
		return p.a
	case 'm':
		return p.m
	case 's':
		return p.s
	case 'x':
		return p.x
	default:
		panic("unexpected attribute")
	}
}

func (p *part) set(n byte, v int) {
	switch n {
	case 'a':
		p.a = v
	case 'm':
		p.m = v
	case 's':
		p.s = v
	case 'x':
		p.x = v
	default:
		panic("unexpected attribute")
	}
}

var (
	inId = idify([]byte("in"))
)

const (
	acceptId = uint32('A')
	rejectId = uint32('R')
)

func parseInput(input []byte) (workflows map[uint32]workflow, parts []part) {
	lines := util.ParseInputLines(input)
	workflows = make(map[uint32]workflow, len(lines))
	li := 0
	var idBytes, ruleBytes, attributeBytes, valueBytes []byte
	for ; li < len(lines); li++ {
		line := lines[li]
		if len(line) == 0 {
			break
		}
		idBytes, line, _ = bytes.Cut(line, []byte("{"))
		line = line[:len(line)-1] // cut ending curly brace
		wf := workflow{rules: make([]rule, 0, 3)}
		for len(line) > 0 {
			ruleBytes, line, _ = bytes.Cut(line, []byte(","))
			compBytes, targetBytes, ok := bytes.Cut(ruleBytes, []byte(":"))
			var r rule
			if !ok {
				r.workflowId = idify(compBytes)
			} else {
				r.attr = compBytes[0]
				r.comp = compBytes[1]
				r.value = util.Btoi(compBytes[2:])
				r.workflowId = idify(targetBytes)
			}
			wf.rules = append(wf.rules, r)
		}
		workflows[idify(idBytes)] = wf
	}
	li++ // skip empty line
	parts = make([]part, 0, len(lines)-li)
	for ; li < len(lines); li++ {
		line := lines[li]
		line = line[1 : len(line)-1] // cut curly braces
		p := part{}
		for len(line) > 0 {
			attributeBytes, line, _ = bytes.Cut(line, []byte(","))
			attributeBytes, valueBytes, _ = bytes.Cut(attributeBytes, []byte("="))
			p.set(attributeBytes[0], util.Btoi(valueBytes))
		}
		parts = append(parts, p)
	}
	return workflows, parts
}

func Run(input []byte) (*util.Result, error) {
	start := time.Now()

	workflows, parts := parseInput(input)

	parse := time.Now()

	part1 := 0
	for _, p := range parts {
		workflowId := inId
		for workflowId != acceptId && workflowId != rejectId {
			wf := workflows[workflowId]
		out:
			for _, r := range wf.rules {
				switch r.comp {
				case '>':
					if p.get(r.attr) > r.value {
						workflowId = r.workflowId
						break out
					}
				case '<':
					if p.get(r.attr) < r.value {
						workflowId = r.workflowId
						break out
					}
				case 0:
					workflowId = r.workflowId
					break out
				default:
					panic("unexpected comparison")
				}
			}
		}
		if workflowId == acceptId {
			part1 += p.x + p.s + p.a + p.m
		}
	}
	part2 := 0
	ranges := make([]partRange, 0, 16)
	ranges = append(ranges, partRange{
		workflowId: inId,
		p1:         part{1, 1, 1, 1},
		p2:         part{4000, 4000, 4000, 4000},
	})
	for len(ranges) > 0 {
		pr := ranges[len(ranges)-1]
		ranges = ranges[:len(ranges)-1]
		if pr.workflowId == acceptId {
			part2 += pr.combinations()
			continue
		} else if pr.workflowId == rejectId {
			continue
		}
		wf := workflows[pr.workflowId]
		for _, r := range wf.rules {
			ranges = append(ranges, pr.apply(r))
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

func idify(idBytes []byte) uint32 {
	id := uint32(0)
	for _, b := range idBytes {
		id = id<<8 | uint32(b)
	}
	return id
}
