package main

import (
	"fmt"
)

func main_1_9_2() {
	var son int

	fmt.Scan(&son)

	yuzlik := son / 100
	onlik := son / 10 % 10
	birlik := son % 10

	if yuzlik == onlik || yuzlik == birlik || onlik == birlik {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
}
