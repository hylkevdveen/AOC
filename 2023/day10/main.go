package main

import (
	p "./pipes"
	"fmt"
)

func main() {
	pipes := p.DrawPipes()
	fmt.Println(Part1(pipes))
	fmt.Println(Part2(pipes))
}

func Part1(pipes *p.Pipes) int {
	// Hardcoded start to make it easier
	pipes.Move(p.West)
	for {
		if pipes.Current == pipes.Start {
			break
		}
		pipes.FollowPipe()
	}
	return (len(pipes.Steps) - 1) / 2
}

func Part2(pipes *p.Pipes) int {
	insideCount := 0
	for _, line := range pipes.MainLoop {
		inside := false
		for _, char := range line {
			if char == 'L' || char == 'J' || char == '|' {
				inside = !inside
			}
			if char == 0 {
				if inside {
					insideCount++
				}
				fmt.Print(".")
			} else {
				fmt.Print(string(char))
			}
		}
		fmt.Print("\n")
	}
	return insideCount
}
