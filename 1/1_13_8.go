package main

import "fmt"

func main_1_13_8() {

	var n int
	fmt.Scan(&n)

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	var sana, minimum int
	minimum = arr[0]

	for i := 0; i < n; i++ {
		if arr[i] < minimum {
			minimum = arr[i]
			sana = 0
		}

		if arr[i] == minimum {
			sana++
		}
	}

	fmt.Println(sana)

}
