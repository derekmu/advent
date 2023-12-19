package day19

import (
	"advent/util"
	"bytes"
	_ "embed"
	"time"
)

type rule struct {
	name       byte
	comp       byte
	value      int
	workflowId uint32
}

type workflow struct {
	name  []byte
	rules []rule
}

type part struct {
	x int
	m int
	a int
	s int
}

func (p *part) get(n byte) int {
	switch n {
	case 'x':
		return p.x
	case 'm':
		return p.m
	case 'a':
		return p.a
	case 's':
		return p.s
	default:
		panic("unexpected attribute")
	}
}

func (p *part) set(n byte, v int) {
	switch n {
	case 'x':
		p.x = v
	case 'm':
		p.m = v
	case 'a':
		p.a = v
	case 's':
		p.s = v
	default:
		panic("unexpected attribute")
	}
}

const (
	acceptId = uint32('A')
	rejectId = uint32('R')
)

var Problem = util.Problem{Year: "2023", Day: "19", Runner: Run, Input: Input}

//go:embed input.txt
var Input []byte

func parseInput(input []byte) (workflows map[uint32]workflow, parts []part) {
	lines := util.ParseInputLines(input)
	workflows = make(map[uint32]workflow, len(lines))
	parts = make([]part, 0, len(lines)/2)
	li := 0
	var idBytes, ruleBytes, attributeBytes, valueBytes []byte
	for ; li < len(lines); li++ {
		line := lines[li]
		if len(line) == 0 {
			break
		}
		idBytes, line, _ = bytes.Cut(line, []byte("{"))
		line = line[:len(line)-1] // cut ending curly brace
		wf := workflow{name: idBytes, rules: make([]rule, 0, 3)}
		for len(line) > 0 {
			ruleBytes, line, _ = bytes.Cut(line, []byte(","))
			compBytes, targetBytes, ok := bytes.Cut(ruleBytes, []byte(":"))
			var r rule
			if !ok {
				r.workflowId = idify(compBytes)
			} else {
				r.name = compBytes[0]
				r.comp = compBytes[1]
				r.value = util.Btoi(compBytes[2:])
				r.workflowId = idify(targetBytes)
			}
			wf.rules = append(wf.rules, r)
		}
		workflows[idify(idBytes)] = wf
	}
	li++ // skip empty line
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
		workflowId := idify([]byte("in"))
		for workflowId != acceptId && workflowId != rejectId {
			wf := workflows[workflowId]
		out:
			for _, r := range wf.rules {
				switch r.comp {
				case '>':
					if p.get(r.name) > r.value {
						workflowId = r.workflowId
						break out
					}
				case '<':
					if p.get(r.name) < r.value {
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

func idify(idBytes []byte) uint32 {
	id := uint32(0)
	for _, b := range idBytes {
		id = id<<8 | uint32(b)
	}
	return id
}
