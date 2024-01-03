package main

import (
	"fmt"
)

func main_1_10_2() {

	var start, finish int
	summa := 0
	fmt.Scan(&start, &finish)

	for start <= finish {
		summa += start
		start++
	}

	fmt.Println(summa)
}
