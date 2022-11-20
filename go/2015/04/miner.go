package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

const initialNumber = 1

func mine(secretKey, requiredPrefix string) int {
	number := initialNumber
	for {
		dataString := fmt.Sprintf("%s%d", secretKey, number)
		data := []byte(dataString)
		md5Sum := md5.Sum(data)
		md5SumHex := hex.EncodeToString(md5Sum[:])
		if strings.HasPrefix(md5SumHex, requiredPrefix) {
			break
		}
		number++
	}
	return number
}
