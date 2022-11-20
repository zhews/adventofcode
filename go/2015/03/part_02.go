package main

import "fmt"

func Part2(directions string) int {
	history := map[string]bool{}
	santaX, santaY, robotX, robotY := initialX, initialY, initialX, initialY
	santasTurn := true
	history[fmt.Sprintf("%dx%d", santaX, santaY)] = true
	for _, direction := range directions {
		var x, y *int
		if santasTurn {
			x = &santaX
			y = &santaY
		} else {
			x = &robotX
			y = &robotY
		}
		santasTurn = !santasTurn
		switch direction {
		case directionNorth:
			*y++
		case directionSouth:
			*y--
		case directionEast:
			*x++
		case directionWest:
			*x--
		}
		history[fmt.Sprintf("%dx%d", *x, *y)] = true
	}
	return len(history)
}
