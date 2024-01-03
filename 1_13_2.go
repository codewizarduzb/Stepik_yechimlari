package main

import "fmt"

func main_1_13_2() {
	var son int
	fmt.Scan(&son)

	son = son%10*100 + son/10%10*10 + son/100

	fmt.Println(son)
}
