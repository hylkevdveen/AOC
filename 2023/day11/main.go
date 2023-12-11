package main

import (
	"../../aoc"
	u "./universe"
	"fmt"
)

const (
	Expansion1 = 2
	Expansion2 = 1_000_000
)

func main() {
	fmt.Println(Part1())
	fmt.Println(Part2())
}

func Part1() int {
	return galaxyDistances(u.ReadGalaxy(Expansion1))
}

func Part2() int {
	return galaxyDistances(u.ReadGalaxy(Expansion2))
}

func galaxyDistances(universe *u.Universe) int {
	total := 0
	for i := 0; i < len(universe.Galaxies)-1; i++ {
		galaxy := universe.Galaxies[i]
		for j := i + 1; j < len(universe.Galaxies); j++ {
			other := universe.Galaxies[j]
			total += aoc.AbsDiffInt(galaxy.X, other.X) + aoc.AbsDiffInt(galaxy.Y, other.Y)
		}
	}
	return total
}
