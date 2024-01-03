package main

import "fmt"

func main_1_12_5() {
	var n, sum int
	fmt.Scan(&n)
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	for i := 0; i < n; i++ {
		if arr[i] > 0 {
			sum++
		}
	}
	fmt.Println(sum)
}
