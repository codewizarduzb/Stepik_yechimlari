package main

import "fmt"

func main_1_13_4() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if a*a+b*b == c*c {
		fmt.Println("Прямоугольный")
	} else {
		fmt.Println("Непрямоугольный")
	}
}
