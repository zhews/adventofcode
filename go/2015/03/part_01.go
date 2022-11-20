package main

import "fmt"

func Part1(directions string) int {
	history := map[string]bool{}
	x, y := initialX, initialY
	history[fmt.Sprintf("%dx%d", x, y)] = true
	for _, direction := range directions {
		switch direction {
		case directionNorth:
			y++
		case directionSouth:
			y--
		case directionEast:
			x++
		case directionWest:
			x--
		}
		history[fmt.Sprintf("%dx%d", x, y)] = true
	}
	return len(history)
}
