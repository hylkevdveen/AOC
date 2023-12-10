package main

import (
	"../../aoc"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	histories := make([][]int, 0)
	for _, line := range aoc.InputLines() {
		numbers := make([]int, 0)
		for _, val := range strings.Split(line, " ") {
			number, _ := strconv.Atoi(val)
			numbers = append(numbers, number)
		}
		histories = append(histories, numbers)
	}
	fmt.Println(Part1(histories))
	fmt.Println(Part2(histories))
}

func Part1(histories [][]int) int {
	total := 0
	for _, history := range histories {
		total += diffs(history) + history[len(history)-1]
	}
	return total
}

func diffs(history []int) int {
	differences := make([]int, 0, len(history)-1)
	for i := 1; i < len(history); i++ {
		differences = append(differences, history[i]-history[i-1])
	}
	if len(differences) == 0 {
		return 0
	}
	return diffs(differences) + differences[len(differences)-1]
}

func Part2(histories [][]int) int {
	total := 0
	for _, history := range histories {
		total += diffs(aoc.ReverseSlice(history)) + history[0]
	}
	return total
}
