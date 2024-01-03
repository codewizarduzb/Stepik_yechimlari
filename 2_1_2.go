package main

import (
	"fmt"
)

func minimumFromFour() int {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	var arr = []int{}
	var minimal_son = a

	arr = append(arr, a, b, c, d)

	for i := 0; i < len(arr); i++ {
		if arr[i] < minimal_son {
			minimal_son = arr[i]
		}
	}

	return minimal_son
}
