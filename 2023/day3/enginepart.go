package main

import "../../aoc"

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
