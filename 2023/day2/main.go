package main

import (
	"../../aoc"
	"fmt"
	"strconv"
	"strings"
)

// ===========
// == Color ==
// ===========

type Color struct {
	Name string
	Max  int
}

var (
	Red   = &Color{Name: "red", Max: 12}
	Green = &Color{Name: "green", Max: 13}
	Blue  = &Color{Name: "blue", Max: 14}
)

var Colors = []*Color{Red, Green, Blue}

func getByName(name string) *Color {
	for _, color := range Colors {
		if color.Name == name {
			return color
		}
	}
	return nil
}

// Main

func main() {
	fmt.Println(Part1())
	fmt.Println(Part2())
}

// Part1 Sum all game IDs of games where the hands drawn could be valid
func Part1() int {
	sum := 0
	for idx, line := range aoc.InputLines() {
		game := line[strings.Index(line, ": ")+2:]
		hands := strings.Split(game, "; ")
		valid := true
		for _, hand := range hands {
			if !handIsValid(hand) {
				valid = false
				break
			}
		}
		if valid {
			sum += idx + 1
		}
	}
	return sum
}

func handIsValid(hand string) bool {
	cubes := strings.Split(hand, ", ")
	for _, cube := range cubes {
		color, count := parseCube(cube)
		if count > color.Max {
			return false
		}
	}
	return true
}

func parseCube(cube string) (*Color, int) {
	spacePos := strings.Index(cube, " ")
	color := getByName(cube[spacePos+1:])
	count, _ := strconv.Atoi(cube[:spacePos])
	return color, count
}

// Part2 Sum the products of the minimum number of red, green and blue cubes needed to have the hands for each game be valid
func Part2() int {
	sum := 0
	for _, line := range aoc.InputLines() {
		game := line[strings.Index(line, ": ")+2:]
		sum += prodMinCubes(game)
	}
	return sum
}

func prodMinCubes(game string) int {
	Red.Max = 0
	Green.Max = 0
	Blue.Max = 0
	hands := strings.Split(game, "; ")
	for _, hand := range hands {
		cubes := strings.Split(hand, ", ")
		for _, cube := range cubes {
			color, count := parseCube(cube)
			if count > color.Max {
				color.Max = count
			}
		}
	}
	return Red.Max * Green.Max * Blue.Max
}
