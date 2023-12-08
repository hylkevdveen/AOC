package network

import (
	"../../../aoc"
	"regexp"
	"strings"
)

type Network struct {
	nodes            map[string][]string
	instructions     []uint8
	InstructionIndex int
	StartNodes       []string
	CurrentNode      string
	Steps            int
}

func ReadNetwork() *Network {
	lines := aoc.InputLines()
	instructionLine := lines[0]
	instructions := make([]uint8, 0, len(instructionLine))
	for _, char := range instructionLine {
		var instruction uint8
		if char == 'L' {
			instruction = 0
		} else {
			instruction = 1
		}
		instructions = append(instructions, instruction)
	}
	nodes := make(map[string][]string)
	startingNodes := make([]string, 0)
	nodeRegexp := regexp.MustCompile("([A-Z]{3}) = \\(([A-Z]{3}), ([A-Z]{3})\\)")
	for _, line := range lines[2:] {
		m := nodeRegexp.FindStringSubmatch(line)
		nodes[m[1]] = []string{m[2], m[3]}
		if strings.HasSuffix(m[1], "A") {
			startingNodes = append(startingNodes, m[1])
		}
	}
	return &Network{
		nodes:            nodes,
		instructions:     instructions,
		InstructionIndex: 0,
		StartNodes:       startingNodes,
		CurrentNode:      "AAA",
		Steps:            0,
	}
}

func (n *Network) StepsToEnd(isEnd func(node string) bool) int {
	for {
		if isEnd(n.CurrentNode) {
			break
		}
		n.CurrentNode = n.nodes[n.CurrentNode][n.instructions[n.InstructionIndex]]
		n.Steps++
		n.InstructionIndex++
		n.InstructionIndex %= len(n.instructions)
	}
	return n.Steps
}
