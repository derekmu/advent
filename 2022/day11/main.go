package day11

import (
	"advent/util"
	"bytes"
	_ "embed"
	"fmt"
	"time"
)

//go:embed input.txt
var Input []byte
var Problem = util.Problem{Year: "2022", Day: "11", Runner: Run, Input: Input}

type op struct {
	name  byte
	value int
}

type monkey struct {
	items1       []int
	items2       []int
	operation    op
	test         int
	throwTrue    int
	throwFalse   int
	inspections1 int
	inspections2 int
}

func Run(input []byte) (util.Result, error) {
	start := time.Now()

	lines := util.ParseInputLines(input)
	monkeys := make([]*monkey, 0, 7)
	var item []byte
	var found bool
	for li := 0; li < len(lines); li += 7 {
		m := &monkey{}
		_, itemsLine, _ := bytes.Cut(lines[li+1], []byte(": "))
		m.items1 = make([]int, 0, (len(itemsLine)+2)/4)
		m.items2 = make([]int, 0, (len(itemsLine)+2)/4)
		for {
			item, itemsLine, found = bytes.Cut(itemsLine, []byte(", "))
			i := util.Btoi(item)
			m.items1 = append(m.items1, i)
			m.items2 = append(m.items2, i)
			if !found {
				break
			}
		}
		operationLine := lines[li+2]
		opi := bytes.LastIndexAny(operationLine, "*+")
		valueLine := operationLine[opi+2:]
		value := -1
		if !bytes.Equal(valueLine, []byte("old")) {
			value = util.Btoi(valueLine)
		}
		m.operation = op{
			name:  operationLine[opi],
			value: value,
		}
		testLine := lines[li+3]
		ti := bytes.LastIndexByte(testLine, 'y')
		m.test = util.Btoi(testLine[ti+2:])
		trueLine := lines[li+4]
		m.throwTrue = int(trueLine[len(trueLine)-1] - '0')
		falseLine := lines[li+5]
		m.throwFalse = int(falseLine[len(falseLine)-1] - '0')
		monkeys = append(monkeys, m)
	}

	parse := time.Now()

	mDivisors := 1
	for _, m := range monkeys {
		mDivisors *= m.test
	}

	for round := 0; round < 10000; round++ {
		for _, m := range monkeys {
			if round < 20 {
				for _, i := range m.items1 {
					v := m.operation.value
					if m.operation.value == -1 {
						v = i
					}
					switch m.operation.name {
					case '+':
						i += v
					case '*':
						i *= v
					default:
						panic(fmt.Sprintf("unknown operation %c", m.operation.name))
					}
					i /= 3
					if i%m.test == 0 {
						monkeys[m.throwTrue].items1 = append(monkeys[m.throwTrue].items1, i)
					} else {
						monkeys[m.throwFalse].items1 = append(monkeys[m.throwFalse].items1, i)
					}
				}
				m.inspections1 += len(m.items1)
				m.items1 = m.items1[:0]
			}

			for _, i := range m.items2 {
				v := m.operation.value
				if m.operation.value == -1 {
					v = i
				}
				switch m.operation.name {
				case '+':
					i += v
				case '*':
					i *= v
				default:
					panic(fmt.Sprintf("unknown operation %c", m.operation.name))
				}
				i = i % mDivisors
				if i%m.test == 0 {
					monkeys[m.throwTrue].items2 = append(monkeys[m.throwTrue].items2, i)
				} else {
					monkeys[m.throwFalse].items2 = append(monkeys[m.throwFalse].items2, i)
				}
			}
			m.inspections2 += len(m.items2)
			m.items2 = m.items2[:0]
		}
	}

	maxInspections := [2]int{}
	for _, m := range monkeys {
		if m.inspections1 > maxInspections[0] {
			maxInspections[1] = maxInspections[0]
			maxInspections[0] = m.inspections1
		} else if m.inspections1 > maxInspections[1] {
			maxInspections[1] = m.inspections1
		}
	}
	part1 := maxInspections[0] * maxInspections[1]

	maxInspections = [2]int{}
	for _, m := range monkeys {
		if m.inspections2 > maxInspections[0] {
			maxInspections[1] = maxInspections[0]
			maxInspections[0] = m.inspections2
		} else if m.inspections2 > maxInspections[1] {
			maxInspections[1] = m.inspections2
		}
	}
	part2 := maxInspections[0] * maxInspections[1]

	end := time.Now()

	return util.Result{
		Part1:     part1,
		Part2:     part2,
		StartTime: start,
		ParseTime: parse,
		EndTime:   end,
	}, nil
}
