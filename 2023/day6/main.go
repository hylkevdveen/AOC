package main

import (
	r "./race"
	"fmt"
)

func main() {
	fmt.Println(Part1(r.ReadRaces()))
	fmt.Println(Part2(r.ReadSingleRace()))
}

// Part1 Multiply the number of ways to win each race
func Part1(races []*r.Race) int {
	winningWays := 1
	for _, race := range races {
		low, high := race.WinningHoldTimes()
		winningWays *= high - low + 1
	}
	return winningWays
}

// Part2 Do the same as Part1 but read the input as a single race
func Part2(race *r.Race) int {
	low, high := race.WinningHoldTimes()
	return high - low + 1
}
