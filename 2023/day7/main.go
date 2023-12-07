package main

import (
	p "./pokerhand"
	"fmt"
)

func main() {
	fmt.Println(Part1())
	fmt.Println(Part2())
}

func Part1() int {
	pokerHands := p.ReadHands(false)
	p.SortHands(pokerHands)
	totalWinnings := 0
	for i, hand := range pokerHands {
		totalWinnings += (i + 1) * hand.Bet
	}
	return totalWinnings
}

func Part2() int {
	pokerHands := p.ReadHands(true)
	p.SortHands(pokerHands)
	totalWinnings := 0
	for i, hand := range pokerHands {
		totalWinnings += (i + 1) * hand.Bet
	}
	return totalWinnings
}
