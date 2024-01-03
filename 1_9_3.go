package main

import (
	"fmt"
)

func main_1_9_3() {

	var num int
	fmt.Scan(&num)

	if num > 0 && num < 10 {
		fmt.Println(num)
	} else if num >= 10 && num < 100 {
		fmt.Println(num / 10)
	} else if num >= 100 && num < 1000 {
		fmt.Println(num / 100)
	} else if num >= 1000 && num < 10000 {
		fmt.Println(num / 1000)
	} else {
		fmt.Println(1)
	}

}
