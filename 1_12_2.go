package main

import (
	"fmt"
)

func main_1_12_2() {
	var n int
	fmt.Scan(&n)

	var array []int

	for n != 0 {
		var son int
		fmt.Scan(&son)
		array = append(array, son)
		n--
	}
	fmt.Println(array[3])
}
