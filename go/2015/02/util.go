package main

import (
	"errors"
	"strconv"
	"strings"
)

const errorInvalidInstruction = "instruction is invalid, it should have 3 dimensions"

func parseInstruction(instruction string) (int, int, int, error) {
	dimensions := strings.Split(instruction, "x")
	if len(dimensions) != 3 {
		return 0, 0, 0, errors.New(errorInvalidInstruction)
	}
	lengthString := dimensions[0]
	length, err := strconv.Atoi(lengthString)
	if err != nil {
		return 0, 0, 0, err
	}
	widthString := dimensions[1]
	width, err := strconv.Atoi(widthString)
	if err != nil {
		return 0, 0, 0, err
	}
	heightString := dimensions[2]
	height, err := strconv.Atoi(heightString)
	if err != nil {
		return 0, 0, 0, err
	}
	return length, width, height, nil
}
