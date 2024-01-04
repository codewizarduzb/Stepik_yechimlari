package main

import "fmt"

func main_1_13_7() {
	var n, sana int
	fmt.Scan(&n)

	for n != 0 {
		var son int
		fmt.Scan(&son)

		if son == 0 {
			sana++
		}
		n--
	}

	fmt.Println(sana)

}
