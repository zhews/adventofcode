package main

import "strings"

func Part02(input string) int {
	text := strings.Split(input, "\n")
	niceStrings := 0
	for _, line := range text {
		repeatingPair := false
		for index := 0; index < len(line)-1; index++ {
			pair := line[index : index+2]
			firstIndex := strings.Index(line, pair)
			lastIndex := strings.LastIndex(line, pair)
			if firstIndex != lastIndex && firstIndex+1 != lastIndex {
				repeatingPair = true
				break
			}
		}
		if !repeatingPair {
			continue
		}
		repeatingRuneSeperated := false
		for index := 0; index < len(line)-2; index++ {
			if line[index] == line[index+2] {
				repeatingRuneSeperated = true
				break
			}
		}
		if !repeatingRuneSeperated {
			continue
		}
		niceStrings++
	}
	return niceStrings
}
