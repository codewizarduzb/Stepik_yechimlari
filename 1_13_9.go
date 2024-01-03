package main

import (
	"fmt"
)

func main_1_13_9() {
	var son int
	fmt.Scan(&son)
	sum := 0

	for son > 0 {
		sum += son % 10
		son /= 10
	}

	fmt.Println(sum/10 + sum%10)

}
