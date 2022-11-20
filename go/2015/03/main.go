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
	solutionPart1 := Part01(input)
	fmt.Printf("Part 1: %d\n", solutionPart1)
	solutionPart2 := Part02(input)
	fmt.Printf("Part 2: %d\n", solutionPart2)
}
