package main

import (
	"strconv"
	"strings"
)

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
