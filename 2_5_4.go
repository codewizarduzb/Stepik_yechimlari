package main

import (
	"fmt"
)

func main_2_5_4() {

	var word string
	fmt.Scan(&word)

	rs := []rune(word)

	for i := range rs {
		if i%2 == 1 {
			fmt.Print(string(rs[i]))
		}
	}
}
