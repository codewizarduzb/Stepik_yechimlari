package main

import (
	"fmt"
)

func main_1_10_4() {

	var son, sana, maksimal_son int
	maksimal_son = son
	sana = 1

	for {
		fmt.Scan(&son)

		if son == 0 {
			break
		}

		if son > maksimal_son {
			maksimal_son = son
			sana = 1
		} else if son == maksimal_son {
			sana++
		}
	}
	fmt.Println(sana)
}
