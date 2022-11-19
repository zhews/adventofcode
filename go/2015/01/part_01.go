package main

func Part1(input string) int {
	floor := initialFloor
	for _, instruction := range input {
		if instruction == instructionUp {
			floor++
		}
		if instruction == instructionDown {
			floor--
		}
	}
	return floor
}
