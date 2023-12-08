package main

import (
	"../../aoc"
	n "./network"
	"fmt"
)

func main() {
	network := n.ReadNetwork()
	fmt.Println(Part1(network))
	fmt.Println(Part2(network))
}

func Part1(network *n.Network) int {
	return network.StepsToEnd(func(node string) bool { return node == "ZZZ" })
}

func Part2(network *n.Network) int {
	ends := make([]int, 0, len(network.StartNodes))
	for _, node := range network.StartNodes {
		network.CurrentNode = node
		network.Steps = 0
		network.InstructionIndex = 0
		ends = append(ends, network.StepsToEnd(func(node string) bool { return node[2] == 'Z' }))
	}
	return aoc.LCM(ends[0], ends[1], ends[2:]...)
}
