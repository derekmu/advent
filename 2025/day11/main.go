package day11

import (
	"advent/util"
	"bytes"
	_ "embed"
	"os"
	"time"

	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
)

//go:embed input.txt
var Input []byte
var Problem = util.Problem{Year: "2025", Day: "11", Runner: Run, Input: Input}

type node struct {
	id                     string
	outputIds              []string
	outputs                []*node
	inputs                 []*node
	paths                  int  // number of paths that reach this node
	visits, expectedVisits int  // how many inputs have been processed and how many are needed before processing this node
	visited                bool // used for tracking which nodes have been queued for marking
	you, fft, dac          bool // marking for being connected to the "you", "fft", and "dac" nodes
}

func parseInput(input []byte) (nodes map[string]*node) {
	lines := util.ParseInputLines(input)
	nodes = make(map[string]*node, len(lines)+1)
	nodes["out"] = &node{
		id: "out",
	}
	for _, line := range lines {
		label, output, _ := bytes.Cut(line, []byte(": "))
		id := string(label)
		outputs := util.ParseInputDelimiter(output, []byte(" "))
		outIds := make([]string, len(outputs))
		for i, o := range outputs {
			outIds[i] = string(o)
		}
		nodes[id] = &node{
			id:        id,
			outputIds: outIds,
		}
	}
	for _, n := range nodes {
		for _, id := range n.outputIds {
			on := nodes[id]
			on.inputs = append(on.inputs, n)
			n.outputs = append(n.outputs, on)
		}
	}
	return nodes
}

func Run(input []byte) (util.Result, error) {
	start := time.Now()

	nodes := parseInput(input)

	parse := time.Now()

	markConnected(nodes, "you", true, false, func(n *node) { n.you = true })
	markConnected(nodes, "fft", true, true, func(n *node) { n.fft = true })
	markConnected(nodes, "dac", true, true, func(n *node) { n.dac = true })
	//updateGraph(nodes, "full_graph", func(n *node) bool { return true })
	//updateGraph(nodes, "you_graph", func(n *node) bool { return n.you })
	//updateGraph(nodes, "fft_dac_graph", func(n *node) bool { return n.fft && n.dac })

	part1 := solve(nodes, "you", "out", func(n *node) bool { return n.you })
	part2 := solve(nodes, "svr", "out", func(n *node) bool { return n.fft && n.dac })

	end := time.Now()

	return util.Result{
		Part1:     part1,
		Part2:     part2,
		StartTime: start,
		ParseTime: parse,
		EndTime:   end,
	}, nil
}

func solve(nodes map[string]*node, start string, end string, f func(n *node) bool) int {
	for _, n := range nodes {
		n.expectedVisits = 0
		n.visits = 0
		n.paths = 0
		if f(n) {
			for _, in := range n.inputs {
				if f(in) {
					n.expectedVisits++
				}
			}
		}
	}

	queue := make([]*node, 0, len(nodes))
	queue = append(queue, nodes[start])
	nodes[start].paths = 1
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		for _, on := range n.outputs {
			if f(on) {
				on.paths += n.paths
				on.visits++
				if on.visits == on.expectedVisits {
					queue = append(queue, on)
				}
			}
		}
	}
	return nodes[end].paths
}

func markConnected(nodes map[string]*node, s string, markOut, markIn bool, f func(n *node)) {
	if markOut {
		clearVisited(nodes)
		markq := make([]*node, 0, len(nodes))
		markq = append(markq, nodes[s])
		for len(markq) > 0 {
			n := markq[0]
			markq = markq[1:]
			f(n)
			for _, on := range n.outputs {
				if !on.visited {
					on.visited = true
					markq = append(markq, on)
				}
			}
		}
	}

	if markIn {
		clearVisited(nodes)
		markq := make([]*node, 0, len(nodes))
		markq = append(markq, nodes[s])
		for len(markq) > 0 {
			n := markq[0]
			markq = markq[1:]
			f(n)
			for _, in := range n.inputs {
				if !in.visited {
					in.visited = true
					markq = append(markq, in)
				}
			}
		}
	}
}

func clearVisited(nodes map[string]*node) {
	for _, n := range nodes {
		n.visited = false
	}
}

func updateGraph(nodes map[string]*node, name string, f func(n *node) bool) {
	g := graph.New(graph.StringHash, graph.Directed())
	for _, n := range nodes {
		if f(n) {
			g.AddVertex(n.id)
		}
	}
	for _, n := range nodes {
		for _, on := range n.outputs {
			g.AddEdge(n.id, on.id)
		}
	}

	file, _ := os.Create("2025/day11/" + name + ".gv")
	_ = draw.DOT(g, file)
	// run this command from the command line to make an SVG to view the graph
	// dot -Tsvg -O graph.gv
}
