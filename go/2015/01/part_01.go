package main

import "errors"

func Part1(input string) (int, error) {
	if len(input) == 0 {
		return 0, errors.New(errorEmptyInstructions)
	}
	floor := initialFloor
	for _, instruction := range input {
		if instruction == instructionUp {
			floor++
		}
		if instruction == instructionDown {
			floor--
		}
	}
	return floor, nil
}
