package main

import (
	"fmt"
)

func main_1_9_4() {
	var son int
	fmt.Scan(&son)

	a1 := son / 100000
	a2 := son / 10000 % 10
	a3 := son / 1000 % 10
	a4 := son / 100 % 10
	a5 := son / 10 % 10
	a6 := son % 10

	if a1+a2+a3 == a4+a5+a6 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
