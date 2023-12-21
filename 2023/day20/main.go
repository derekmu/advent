package day20

import (
	"advent/util"
	"bytes"
	_ "embed"
	"time"
)

//go:embed input.txt
var Input []byte
var Problem = util.Problem{Year: "2023", Day: "20", Runner: Run, Input: Input}

var (
	buttonId      = idify([]byte("button"))
	broadcasterId = idify([]byte("broadcaster"))
	rxId          = idify([]byte("rx"))
	gfId          = uint32(0)
)

type pulse struct {
	fromId uint32
	toId   uint32
	high   bool
}

type module struct {
	mtype byte
	id    uint32
	// whether a flip-flop module is on
	state   bool
	targets []uint32
	// set of connected modules with a low pulse for conjunction modules
	inputs  map[uint32]bool
	gfFlips map[uint32][]int
}

func (m *module) pulse(pulseQueue []pulse, press int) []pulse {
	output := pulseQueue[0].high
	switch m.mtype {
	case '%':
		if pulseQueue[0].high {
			// If a flip-flop module receives a high pulse, it is ignored and nothing happens.
			return pulseQueue
		} else {
			// If a flip-flop module receives a low pulse, it flips between on and off.
			// If it was off, it turns on and sends a high pulse.
			// If it was on, it turns off and sends a low pulse.
			m.state = !m.state
			output = m.state
		}
	case '&':
		if m.id == gfId {
			if m.gfFlips == nil {
				m.gfFlips = make(map[uint32][]int, len(m.inputs))
				for id := range m.inputs {
					m.gfFlips[id] = []int{-1, -1}
				}
			}
			if pulseQueue[0].high {
				for i, v := range m.gfFlips[pulseQueue[0].fromId] {
					if v == -1 {
						m.gfFlips[pulseQueue[0].fromId][i] = press
						break
					}
				}
			}
		}
		// When a pulse is received, the conjunction module first updates its memory for that input.
		if pulseQueue[0].high {
			delete(m.inputs, pulseQueue[0].fromId)
		} else {
			m.inputs[pulseQueue[0].fromId] = false
		}
		// If it remembers high pulses for all inputs, it sends a low pulse; otherwise, it sends a high pulse.
		output = len(m.inputs) > 0
	}
	for _, t := range m.targets {
		pulseQueue = append(pulseQueue, pulse{
			fromId: m.id,
			toId:   t,
			high:   output,
		})
	}
	return pulseQueue
}

func parseInput(input []byte) (moduleMap map[uint32]*module) {
	lines := util.ParseInputLines(input)
	moduleMap = make(map[uint32]*module, len(lines))
	for _, line := range lines {
		name, targetsCombined, _ := bytes.Cut(line, []byte(" -> "))
		mtype := name[0]
		if mtype == '%' || mtype == '&' {
			name = name[1:]
		}
		targetNames := bytes.Split(targetsCombined, []byte(", "))
		targets := make([]uint32, len(targetNames))
		for i, n := range targetNames {
			targets[i] = idify(n)
		}
		m := &module{
			mtype:   mtype,
			id:      idify(name),
			targets: targets,
			inputs:  make(map[uint32]bool, 3),
		}
		moduleMap[m.id] = m
	}
	for n, m := range moduleMap {
		for _, t := range m.targets {
			m2, ok := moduleMap[t]
			if !ok {
				m2 = &module{
					mtype:  ' ',
					id:     t,
					inputs: make(map[uint32]bool, 3),
				}
				moduleMap[t] = m2
			}
			// Conjunction modules remember the type of the most recent pulse received from each of
			// their connected input modules; they initially default to remembering a low pulse for each input.
			m2.inputs[n] = false
		}
	}
	return moduleMap
}

func Run(input []byte) (*util.Result, error) {
	start := time.Now()

	moduleMap := parseInput(input)

	parse := time.Now()

	low := 0
	high := 0
	part2 := -1
	// no sample for part 2
	if m, ok := moduleMap[rxId]; !ok {
		part2 = 0
	} else {
		for k := range m.inputs {
			gfId = k
		}
	}
	reuseQueue := make([]pulse, 0, 100)
	for press := 0; press < 1000 || part2 == -1; press++ {
		pulseQueue := reuseQueue[:0]
		pulseQueue = append(pulseQueue, pulse{fromId: buttonId, toId: broadcasterId, high: false})
		for len(pulseQueue) > 0 {
			if press < 1000 {
				if pulseQueue[0].high {
					high++
				} else {
					low++
				}
			}
			if m, ok := moduleMap[pulseQueue[0].toId]; ok {
				pulseQueue = m.pulse(pulseQueue, press)
				if pulseQueue[0].toId == gfId {
					allDone := true
					for _, v := range m.gfFlips {
						for _, v2 := range v {
							if v2 == -1 {
								allDone = false
								break
							}
						}
					}
					if allDone {
						numbers := make([]int, len(m.gfFlips))
						i := 0
						for _, v := range m.gfFlips {
							numbers[i] = v[1] - v[0]
							i++
						}
						part2 = util.FindLcm(numbers)
					}
				}
			}
			pulseQueue = pulseQueue[1:]
		}
	}
	part1 := low * high

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
	for i, b := range idBytes {
		id = id<<8 | uint32(b)
		if i == 3 {
			break
		}
	}
	return id
}
