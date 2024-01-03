package main

import "fmt"

func main_1_12_3() {
	array := [5]int{}
	var a int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		array[i] = a
	}

	maximal := array[0]
	for i := 0; i < 5; i++ {
		if array[i] > maximal {
			maximal = array[i]
		}
	}
	fmt.Println(maximal)
}
