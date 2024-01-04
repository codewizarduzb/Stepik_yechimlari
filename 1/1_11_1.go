package main

import (
	"fmt"
)

func main_1_11_1() {

	var son float64

	fmt.Scan(&son)

	if son <= 0 {
		fmt.Printf("число %.2f не подходит", son)
	} else if son > 10000 {
		fmt.Printf("%e", son)
	} else {
		fmt.Printf("%.4f", son*son)
	}
}
