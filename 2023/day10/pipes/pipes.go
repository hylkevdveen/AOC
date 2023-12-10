package pipes

import "../../../aoc"

type Coordinate struct {
	X int
	Y int
}

type Pipes struct {
	System    [][]rune
	MainLoop  [][]rune
	Start     Coordinate
	Current   Coordinate
	Steps     []Coordinate
	Direction int
}

const (
	North = iota
	East
	South
	West
)

func DrawPipes() *Pipes {
	var grid [][]rune
	var mainLoop [][]rune
	start := Coordinate{}
	for y, line := range aoc.InputLines() {
		row := []rune(line)
		for x, val := range row {
			if val == 'S' {
				start = Coordinate{X: x, Y: y}
			}
		}
		grid = append(grid, row)
		mainLoop = append(mainLoop, make([]rune, len(row)))
	}
	return &Pipes{
		System:    grid,
		MainLoop:  mainLoop,
		Start:     start,
		Current:   start,
		Steps:     []Coordinate{start},
		Direction: -1,
	}
}

func (p *Pipes) Move(direction int) {
	newPos := Coordinate{p.Current.X, p.Current.Y}
	switch direction {
	case North:
		newPos.Y--
	case East:
		newPos.X++
	case South:
		newPos.Y++
	case West:
		newPos.X--
	}
	p.Steps = append(p.Steps, newPos)
	p.Current = newPos
	p.Direction = direction
	p.MainLoop[newPos.Y][newPos.X] = p.System[newPos.Y][newPos.X]
}

func (p *Pipes) FollowPipe() {
	switch p.System[p.Current.Y][p.Current.X] {
	case '-', '|':
		p.Move(p.Direction)
	case 'L':
		if p.Direction == West {
			p.Move(North)
		} else {
			p.Move(East)
		}
	case 'F':
		if p.Direction == North {
			p.Move(East)
		} else {
			p.Move(South)
		}
	case '7':
		if p.Direction == East {
			p.Move(South)
		} else {
			p.Move(West)
		}
	case 'J':
		if p.Direction == South {
			p.Move(West)
		} else {
			p.Move(North)
		}
	}
}
