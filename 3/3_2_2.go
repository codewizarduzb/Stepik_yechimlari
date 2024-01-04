package main

import (
	"strconv"
	"unicode"
)

func adding(s1, s2 string) int64 {

	var cleanedS1 string
	for _, char := range s1 {
		if unicode.IsDigit(char) {
			cleanedS1 += string(char)
		}
	}
	num1, _ := strconv.ParseInt(cleanedS1, 10, 64)

	var cleanedS2 string
	for _, char := range s2 {
		if unicode.IsDigit(char) {
			cleanedS2 += string(char)
		}
	}
	num2, _ := strconv.ParseInt(cleanedS2, 10, 64)

	return num1 + num2
}
