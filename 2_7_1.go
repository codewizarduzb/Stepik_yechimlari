package main

import (
	"fmt"
	"math"
)

func main_2_7_1() {
	var a, b float64

	fmt.Scan(&a, &b)

	fmt.Println(gipotenuza(a, b))
}

func gipotenuza(a, b float64) float64 {
	return math.Sqrt(a*a + b*b)
}
