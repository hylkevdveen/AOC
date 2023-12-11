package universe

import (
	"../../../aoc"
	"strings"
)

type Coordinate struct {
	X int
	Y int
}

type Universe struct {
	Galaxies []Coordinate
}

func ReadGalaxy(expansionSize int) *Universe {
	lines := aoc.InputLines()
	emptyColumns := map[int]bool{}
	for i := len(lines[0]) - 1; i >= 0; i-- {
		empty := true
		for _, line := range lines {
			if line[i] == '#' {
				empty = false
				break
			}
		}
		if empty {
			emptyColumns[i] = true
		}
	}

	galaxies := make([]Coordinate, 0)
	y := 0
	for i, line := range lines {
		if !strings.ContainsRune(line, '#') {
			y += expansionSize
			continue
		}
		x := 0
		for j := 0; j < len(line); j++ {
			if _, ok := emptyColumns[j]; ok {
				x += expansionSize
				continue
			}
			if lines[i][j] == '#' {
				galaxies = append(galaxies, Coordinate{x, y})
			}
			x++
		}
		y++
	}
	return &Universe{
		Galaxies: galaxies,
	}
}
