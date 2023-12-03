package main

import (
	"../../aoc"
	"fmt"
	"strconv"
	"unicode"
)

// ==========
// == Grid ==
// ==========

const Gear = rune(42)
const Period = rune(46)

type EngineSchematics struct {
	grid  [][]rune
	parts []*EnginePart
	maxX  int
	maxY  int
}

type EnginePart struct {
	value  int
	y      int
	xStart int
	xEnd   int
}

func (p *EnginePart) IsAdjacent(x int, y int) bool {
	return aoc.AbsDiffInt(y, p.y) <= 1 &&
		((p.xStart <= x && x <= p.xEnd) ||
			p.xStart-x == 1 ||
			x-p.xEnd == 1)
}

func ReadEngineSchematics() *EngineSchematics {
	var grid [][]rune
	for _, line := range aoc.InputLines() {
		grid = append(grid, []rune(line))
	}
	return &EngineSchematics{
		grid: grid,
		maxX: len(grid[0]) - 1,
		maxY: len(grid) - 1,
	}
}

func (e *EngineSchematics) BuildEnginePart(y int, xStart int, xEnd int) *EnginePart {
	enginePart := EnginePart{
		y:      y,
		xStart: xStart,
		xEnd:   xEnd,
	}
	if !e.IsValidPartNumber(enginePart) {
		return nil
	}
	partNumber, _ := strconv.Atoi(string(e.grid[y][xStart : xEnd+1]))
	enginePart.value = partNumber
	e.parts = append(e.parts, &enginePart)
	return &enginePart
}

func (e *EngineSchematics) IsValidX(x int) bool {
	return x >= 0 && x <= e.maxX
}

func (e *EngineSchematics) IsValidY(y int) bool {
	return y >= 0 && y <= e.maxY
}

// IsValidPartNumber returns whether the given engine part is surrounded by a special character
func (e *EngineSchematics) IsValidPartNumber(part EnginePart) bool {
	if !e.IsValidY(part.y) || !e.IsValidX(part.xStart) || !e.IsValidX(part.xEnd) {
		return false
	}
	for i := part.y - 1; i <= part.y+1; i++ {
		if !e.IsValidY(i) {
			continue
		}
		for j := part.xStart - 1; j <= part.xEnd+1; j++ {
			if !e.IsValidX(j) {
				continue
			}
			if !(unicode.IsDigit(e.grid[i][j]) || e.grid[i][j] == Period) {
				return true
			}
		}
	}
	return false
}

// AdjacentParts returns all EngineParts that are adjacent to the given coordinate
func (e *EngineSchematics) AdjacentParts(x int, y int) []*EnginePart {
	var adjacentParts []*EnginePart
	for _, part := range e.parts {
		if part.IsAdjacent(x, y) {
			adjacentParts = append(adjacentParts, part)
		}
	}
	return adjacentParts
}

// CalculateGearRatio calculates the product of the values of the AdjacentParts to the given coordinate
func (e *EngineSchematics) CalculateGearRatio(x int, y int) int {
	gearRatio := 0
	adjacentParts := e.AdjacentParts(x, y)
	if len(adjacentParts) > 1 {
		gearRatio = 1
		for _, part := range adjacentParts {
			gearRatio *= part.value
		}
	}
	return gearRatio
}

// Main

func main() {
	engineSchematics := ReadEngineSchematics()
	fmt.Println(Part1(engineSchematics))
	fmt.Println(Part2(engineSchematics))
}

// Part1 Sum the values of the engine parts that are adjacent to a symbol
func Part1(engineSchematics *EngineSchematics) int {
	partNumbers := 0
	for row, line := range engineSchematics.grid {
		numberStart := -1
		numberEnd := -1
		for pos := 0; pos < len(line); pos++ {
			if unicode.IsDigit(line[pos]) {
				if numberStart == -1 {
					numberStart = pos
				}
				numberEnd = pos
			} else {
				if numberStart != -1 {
					part := engineSchematics.BuildEnginePart(row, numberStart, numberEnd)
					if part != nil {
						fmt.Printf("%d is a valid part (row %d, %d-%d)\n", part.value, part.y, part.xStart, part.xEnd)
						partNumbers += part.value
					}
					numberStart = -1
				}
			}
		}
		if numberStart != -1 {
			part := engineSchematics.BuildEnginePart(row, numberStart, numberEnd)
			if part != nil {
				fmt.Printf("%d is a valid part (row %d, %d-%d)\n", part.value, part.y, part.xStart, part.xEnd)
				partNumbers += part.value
			}
		}
	}
	return partNumbers
}

// Part2 Sum the gear ratio (product of engine parts) of all gears (asterisks) that connect multiple engine parts
// Uses []EnginePart of Part1
func Part2(engineSchematics *EngineSchematics) int {
	gearRatios := 0
	for row, line := range engineSchematics.grid {
		for col, value := range line {
			if value == Gear {
				gearRatios += engineSchematics.CalculateGearRatio(col, row)
			}
		}
	}
	return gearRatios
}
