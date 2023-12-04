package main

import "fmt"

const Gear = rune(42)

func main() {
	engineSchematics := ReadEngineSchematics()
	engineSchematics.BuildEngineParts()
	fmt.Println(Part1(engineSchematics))
	fmt.Println(Part2(engineSchematics))
}

// Part1 Sum the values of the engine parts that are adjacent to a symbol
func Part1(engineSchematics *EngineSchematics) int {
	partNumbers := 0
	for _, part := range engineSchematics.parts {
		partNumbers += part.value
	}
	return partNumbers
}

// Part2 Sum the gear ratio (product of engine parts) of all gears (asterisks) that connect multiple engine parts
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
