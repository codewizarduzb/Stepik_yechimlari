package main

import (
	"fmt"
)

func main_1_10_8() {
	var num1, num2 int
	fmt.Scan(&num1, &num2)

	result := ""
	for num1 > 0 {
		digit := num1 % 10
		if containsDigit(num2, digit) {
			result = fmt.Sprintf("%d %s", digit, result)
		}
		num1 /= 10
	}

	fmt.Println(result)
}

func containsDigit(num, digit int) bool {
	for num > 0 {
		if num%10 == digit {
			return true
		}
		num /= 10
	}
	return false
}
