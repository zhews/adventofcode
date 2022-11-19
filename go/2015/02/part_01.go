package main

import (
	"strings"
)

const errorInvalidInstruction = "instruction is invalid, it should have 3 dimensions"

func Part1(input string) (int, error) {
	instructions := strings.Split(input, "\n")
	paperAmount := 0
	for _, instruction := range instructions {
		l, w, h, err := parseInstruction(instruction)
		if err != nil {
			return 0, err
		}
		paperAmountForPresent := calculatePaperAmountForPresent(l, w, h)
		paperAmount += paperAmountForPresent
	}
	return paperAmount, nil
}

func calculatePaperAmountForPresent(length, width, height int) int {
	sideLengthWidth := length * width
	sideWidthHeight := width * height
	sideHeightLength := height * length
	box := 2*sideLengthWidth + 2*sideWidthHeight + 2*sideHeightLength
	slack, _ := minInt(sideLengthWidth, sideWidthHeight, sideHeightLength) // error can be ignored because we know that we are not passing an empty array
	return box + slack
}
