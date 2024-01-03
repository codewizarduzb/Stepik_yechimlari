package main

import (
	"fmt"
)

func main_1_10_3() {

	var n int
	sum := 0

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var son int

		fmt.Scan(&son)
		if son%8 == 0 && son > 9 && son < 99 {
			sum += son
		}
	}

	fmt.Println(sum)
}
