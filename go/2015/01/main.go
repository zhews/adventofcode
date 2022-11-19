package main

import (
	"fmt"
	"os"
)

const inputFileName = "input.txt"

func main() {
	inputBytes, err := os.ReadFile(inputFileName)
	if err != nil {
		panic(err)
	}
	input := string(inputBytes)
	solutionPart1 := Part1(input)
	fmt.Printf("Part 1: %d\n", solutionPart1)
	solutionPart2, err := Part2(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part 2: %d\n", solutionPart2)
}
