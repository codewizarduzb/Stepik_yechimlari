package main

import (
	"fmt"
)

func main_2_7_3() {

	var son string
	fmt.Scan(&son)
	maximal := son[0]

	for i := range son {
		if son[i] > maximal {
			maximal = son[i]
		}
	}
	fmt.Println(string(maximal))
}
