package main

import (
	"../../aoc"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Race struct {
	Time     float64
	Distance float64
}

func ReadRaces() []*Race {
	var races []*Race
	lines := aoc.InputLines()
	var times []float64
	for _, val := range strings.Split(strings.TrimPrefix(lines[0], "Time: "), " ") {
		if val == "" {
			continue
		}
		time, _ := strconv.Atoi(val)
		times = append(times, float64(time))
	}
	var distances []float64
	for _, val := range strings.Split(strings.TrimPrefix(lines[1], "Distance: "), " ") {
		if val == "" {
			continue
		}
		distance, _ := strconv.Atoi(val)
		distances = append(distances, float64(distance))
	}
	for i := 0; i < len(times); i++ {
		races = append(races, &Race{
			Time:     times[i],
			Distance: distances[i],
		})
	}
	return races
}

func ReadSingleRace() *Race {
	lines := aoc.InputLines()
	time, _ := strconv.Atoi(strings.Replace(strings.TrimPrefix(lines[0], "Time: "), " ", "", -1))
	distance, _ := strconv.Atoi(strings.Replace(strings.TrimPrefix(lines[1], "Distance: "), " ", "", -1))
	return &Race{
		Time:     float64(time),
		Distance: float64(distance),
	}
}

// WinningHoldTimes calculates the lower and upper bound of hold times that beat the Distance in the Time allowed
// Find all hold times x where (Time - x) * x > Distance
// Solve for (Time - x) * x = Distance
// Time*x - x^2 - Distance = 0
// -x^2 + Time*x - Distance = 0
// -(-x^2 + Time*x - Distance) = 0
// x^2 - Time*x + Distance = 0
// Plug into quadratic formula with:
// a = 1
// b = -Time
// c = Distance
// Lower bound: (Time - sqrt(Time^2 - 4*Distance)) / 2 -> round up to integer
// Upper bound: (Time + sqrt(Time^2 - 4*Distance)) / 2 -> round down to integer
func (r *Race) WinningHoldTimes() (int, int) {
	lowerBound := math.Ceil((r.Time - math.Sqrt(r.Time*r.Time-4*r.Distance)) / 2)
	upperBound := math.Floor((r.Time + math.Sqrt(r.Time*r.Time-4*r.Distance)) / 2)
	return int(lowerBound), int(upperBound)
}

func main() {
	fmt.Println(Part1(ReadRaces()))
	fmt.Println(Part2(ReadSingleRace()))
}

// Part1 Multiply the number of ways to win each race
func Part1(races []*Race) int {
	winningWays := 1
	for _, race := range races {
		low, high := race.WinningHoldTimes()
		winningWays *= high - low + 1
	}
	return winningWays
}

// Part2 Do the same as Part1 but read as a single race
func Part2(race *Race) int {
	low, high := race.WinningHoldTimes()
	return high - low + 1
}
