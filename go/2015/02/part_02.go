package main

import "strings"

func Part2(input string) (int, error) {
	instructions := strings.Split(input, "\n")
	ribbonAmount := 0
	for _, instruction := range instructions {
		length, width, height, err := parseInstruction(instruction)
		if err != nil {
			return 0, err
		}
		minPresentPerimeter := calculateMinPresentPerimeter(length, width, height)
		presentVolume := calculatePresentVolume(length, width, height)
		ribbonAmountForPresent := minPresentPerimeter + presentVolume
		ribbonAmount += ribbonAmountForPresent
	}
	return ribbonAmount, nil
}

func calculateMinPresentPerimeter(length, width, height int) int {
	perimeterLengthWidth := 2*length + 2*width
	perimeterWidthHeight := 2*width + 2*height
	perimeterHeightLength := 2*height + 2*length
	min, _ := minInt(perimeterLengthWidth, perimeterWidthHeight, perimeterHeightLength) // error can be ignored because we know that we are not passing an empty array
	return min
}

func calculatePresentVolume(length, width, height int) int {
	return length * width * height
}
