package main

import (
	"../../aoc"
	"./network"
	"fmt"
)

func main() {
	n := network.ReadNetwork()
	fmt.Println(Part1(n))
	fmt.Println(Part2(n))
}

func Part1(n *network.Network) int {
	return n.StepsToEnd(func(node string) bool { return node == "ZZZ" })
}

func Part2(n *network.Network) int {
	ends := make([]int, 0, len(n.StartNodes))
	for _, node := range n.StartNodes {
		n.CurrentNode = node
		n.Steps = 0
		n.InstructionIndex = 0
		ends = append(ends, n.StepsToEnd(func(node string) bool { return node[2] == 'Z' }))
	}
	return aoc.LCM(ends[0], ends[1], ends[2:]...)
}
