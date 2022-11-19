package main

import "errors"

const (
	errorDidNotEnterBasement = "input never let santa enter the basement"
)

func Part2(instructions string) (int, error) {
	if len(instructions) == 0 {
		return 0, errors.New(errorEmptyInstructions)
	}
	floor := initialFloor
	for index, instruction := range instructions {
		if instruction == instructionUp {
			floor++
		}
		if instruction == instructionDown {
			floor--
		}
		if floor == -1 {
			instructionPosition := index + 1
			return instructionPosition, nil
		}
	}
	return 0, errors.New(errorDidNotEnterBasement)
}
