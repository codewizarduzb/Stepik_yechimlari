package main

import "fmt"

func main_1_9() {
	var son int
	fmt.Scan(&son)

	if son > 0 {
		fmt.Println("Число положительное")
	} else if son < 0 {
		fmt.Println("Число отрицательное")
	} else {
		fmt.Println("Ноль")
	}

}
