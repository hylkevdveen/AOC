package main

import (
	"../../aoc"
	"strconv"
	"unicode"
)

const Period = rune(46)

type EngineSchematics struct {
	grid  [][]rune
	parts []*EnginePart
	maxX  int
	maxY  int
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

func (e *EngineSchematics) BuildEngineParts() {
	for row, line := range e.grid {
		partStart := -1
		partEnd := -1
		for pos := 0; pos < len(line); pos++ {
			if unicode.IsDigit(line[pos]) {
				if partStart == -1 {
					partStart = pos
				}
				partEnd = pos
			} else {
				if partStart != -1 {
					e.BuildEnginePart(row, partStart, partEnd)
					partStart = -1
				}
			}
		}
		if partStart != -1 {
			e.BuildEnginePart(row, partStart, partEnd)
		}
	}
}

func (e *EngineSchematics) BuildEnginePart(y int, xStart int, xEnd int) *EnginePart {
	enginePart := EnginePart{
		y:      y,
		xStart: xStart,
		xEnd:   xEnd,
	}
	if !e.IsValidPart(enginePart) {
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

func (e *EngineSchematics) IsSpecialChar(c rune) bool {
	return !(unicode.IsDigit(c) || c == Period)
}

// IsValidPart returns whether the given engine part is surrounded by a special character
func (e *EngineSchematics) IsValidPart(part EnginePart) bool {
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
			if e.IsSpecialChar(e.grid[i][j]) {
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
