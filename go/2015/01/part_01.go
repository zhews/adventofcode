package main

func Part1(input string) int {
	floor := 0
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
