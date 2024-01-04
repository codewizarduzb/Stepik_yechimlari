package main

import "fmt"

func main_2_6_1() {
	// package main уже объявлен, нужные импорты уже есть
	var a, b int
	_, err := fmt.Scan(&a, &b)

	if err != nil && b == 0 {
		fmt.Println("ошибка")
	} else {
		//	result, err := divide(a, b)

		if err != nil {
			fmt.Println("ошибка")
		} else {
			//	fmt.Println(result)
		}
	}
}
