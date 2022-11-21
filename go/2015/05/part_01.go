package main

import (
	"regexp"
	"strings"
)

const (
	minVowels             = 3
	maxBlacklistedStrings = 0
	findAllAmount         = -1
)

var regexVowels = regexp.MustCompile(`([aeiou])`)
var regexBlacklist = regexp.MustCompile("(ab|cd|pq|xy)")

func Part01(input string) int {
	text := strings.Split(input, "\n")
	niceStrings := 0
	for _, line := range text {
		vowels := regexVowels.FindAllString(line, findAllAmount)
		if len(vowels) < minVowels {
			continue
		}
		repeatingRunes := false
		for index := 0; index < len(line)-1; index++ {
			if line[index] == line[index+1] {
				repeatingRunes = true
				break
			}
		}
		if !repeatingRunes {
			continue
		}
		blacklistedStrings := regexBlacklist.FindAllString(line, findAllAmount)
		if len(blacklistedStrings) > maxBlacklistedStrings {
			continue
		}
		niceStrings++
	}
	return niceStrings
}
