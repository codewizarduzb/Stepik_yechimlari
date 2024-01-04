package main

import "fmt"

func main_2_7_4() {
	var son int
	var arr = []int{}
	fmt.Scan(&son)

	for son != 0 {
		arr = append(arr, ((son % 10) * (son % 10)))
		son /= 10
	}

	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Print(arr[i])
	}

}
