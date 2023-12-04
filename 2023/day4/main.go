package main

import (
	"../../aoc"
	"fmt"
	"strconv"
	"strings"
)

// =================
// == ScratchCard ==
// =================

type ScratchCard struct {
	winningNumbers  []int
	cardNumbers     []int
	matchingNumbers int
	copies          int
}

func ReadScratchCard(line string) *ScratchCard {
	scratchCard := line[strings.Index(line, ": ")+2:]
	winningNumbers := readNumbers(scratchCard[:strings.Index(scratchCard, " | ")])
	cardNumbers := readNumbers(scratchCard[strings.Index(scratchCard, " | ")+3:])
	winCount := matchingNumbers(winningNumbers, cardNumbers)
	return &ScratchCard{
		winningNumbers:  winningNumbers,
		cardNumbers:     cardNumbers,
		matchingNumbers: winCount,
		copies:          1,
	}
}

func readNumbers(numberString string) []int {
	var numbers []int
	numberSlice := strings.Split(numberString, " ")
	for _, number := range numberSlice {
		if number != "" {
			num, _ := strconv.Atoi(number)
			numbers = append(numbers, num)
		}
	}
	return numbers
}

func matchingNumbers(winners []int, numbers []int) int {
	matches := 0
	for _, winner := range winners {
		for _, number := range numbers {
			if winner == number {
				matches++
				break
			}
		}
	}
	return matches
}

// ==========
// == Main ==
// ==========

func main() {
	var scratchCards []*ScratchCard
	for _, line := range aoc.InputLines() {
		scratchCards = append(scratchCards, ReadScratchCard(line))
	}
	fmt.Println(Part1(scratchCards))
	fmt.Println(Part2(scratchCards))
}

func Part1(scratchCards []*ScratchCard) int {
	value := 0
	for _, scratchCard := range scratchCards {
		if scratchCard.matchingNumbers > 0 {
			value += 1 << (scratchCard.matchingNumbers - 1)
		}
	}
	return value
}

func Part2(scratchCards []*ScratchCard) int {
	total := 0
	for idx, scratchCard := range scratchCards {
		for repeat := 0; repeat < scratchCard.copies; repeat++ {
			total++
			if scratchCard.matchingNumbers > 0 {
				for update := 1; update <= scratchCard.matchingNumbers; update++ {
					scratchCards[idx+update].copies++
				}
			}
		}
	}
	return total
}
