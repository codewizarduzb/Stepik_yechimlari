package main

import (
	"fmt"
)

func main_1_10_7() {

	var x, p, y float64
	yil := 0

	fmt.Scan(&x)
	fmt.Scan(&p)
	fmt.Scan(&y)

	for x < y {
		x *= (1 + p/100)
		yil++
	}

	fmt.Println(yil)
}
