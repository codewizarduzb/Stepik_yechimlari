package main

import "fmt"

func main_1_12_1() {
	var workArray [10]uint8

	for i := 0; i < 10; i++ {
		fmt.Scan(&workArray[i])
	}

	for i := 0; i < 3; i++ {
		var a, b int
		fmt.Scan(&a, &b)

		workArray[a], workArray[b] = workArray[b], workArray[a]
	}

	for i := 0; i < 10; i++ {
		fmt.Print(workArray[i], " ")
	}
}
