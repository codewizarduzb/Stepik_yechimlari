package main

import (
	"fmt"
)

func main_1_10_6() {

	var son int

	for {
		fmt.Scan(&son)

		if son > 100 {
			break
		}

		if son < 10 {
			continue
		}

		fmt.Println(son)
	}
}
