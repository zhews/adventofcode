package main

import "errors"

const errorEmptyInput = "cannot compute min of empty array"

func minInt(digits ...int) (int, error) {
	if len(digits) == 0 {
		return 0, errors.New(errorEmptyInput)
	}
	var min int
	for index, digit := range digits {
		if index == 0 || digit < min {
			min = digit
		}
	}
	return min, nil
}
