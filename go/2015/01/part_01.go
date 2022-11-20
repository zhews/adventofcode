package main

import "errors"

func Part01(instructions string) (int, error) {
	if len(instructions) == 0 {
		return 0, errors.New(errorEmptyInstructions)
	}
	floor := initialFloor
	for _, instruction := range instructions {
		if instruction == instructionUp {
			floor++
		}
		if instruction == instructionDown {
			floor--
		}
	}
	return floor, nil
}
