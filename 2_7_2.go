package main

import (
	"fmt"
)

func main_2_7_2() {

	var word string

	fmt.Scan(&word)

	for i := range word {
		if i == len(word)-1 {
			fmt.Print(string(word[i]))
		} else {
			fmt.Print(string(word[i]) + "*")
		}

	}
}
