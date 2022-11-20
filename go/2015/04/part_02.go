package main

const requiredPrefixPart02 = "000000"

func Part02(secretKey string) int {
	return mine(secretKey, requiredPrefixPart02)
}
