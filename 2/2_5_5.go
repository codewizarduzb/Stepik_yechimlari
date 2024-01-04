package main

import (
	"fmt"
	"strings"
)

func main_2_5_5() {

	var word string
	fmt.Scan(&word)

	rs := []rune(word)

	for i := range rs {
		if strings.Count(word, string(rs[i])) < 2 {
			fmt.Print(string(rs[i]))
		}
	}
}
