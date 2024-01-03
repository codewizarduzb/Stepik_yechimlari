package main

import (
	"fmt"
)

func main_1_13_14() {

	var arr = []int{}

	var son, n int

	fmt.Scan(&son, &n)

	for son != 0 {
		arr = append(arr, son % 10)
		son /= 10
	}

	for i := len(arr) - 1; i >= 0; i--{
		if arr[i] != n{
			fmt.Print(arr[i])
		}
	}
}