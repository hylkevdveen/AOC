package main

import (
	"../../aoc"
	"fmt"
)

func main() {
	var scratchCards []*ScratchCard
	for _, line := range aoc.InputLines() {
		scratchCards = append(scratchCards, ReadScratchCard(line))
	}
	fmt.Println(Part1(scratchCards))
	fmt.Println(Part2(scratchCards))
}

// Part1 Calculate the value of all your winning scratchcards
func Part1(scratchCards []*ScratchCard) int {
	value := 0
	for _, scratchCard := range scratchCards {
		if scratchCard.MatchingNumbers > 0 {
			value += 1 << (scratchCard.MatchingNumbers - 1)
		}
	}
	return value
}

// Part2 Calculate the number of scratchcards you get if each winner gives you copies
func Part2(scratchCards []*ScratchCard) int {
	total := 0
	for idx, scratchCard := range scratchCards {
		for repeat := 0; repeat < scratchCard.Copies; repeat++ {
			total++
			if scratchCard.MatchingNumbers > 0 {
				for update := 1; update <= scratchCard.MatchingNumbers; update++ {
					scratchCards[idx+update].Copies++
				}
			}
		}
	}
	return total
}
