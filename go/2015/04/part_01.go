package main

const requiredPrefixPart01 = "00000"

func Part01(secretKey string) int {
	return mine(secretKey, requiredPrefixPart01)
}
