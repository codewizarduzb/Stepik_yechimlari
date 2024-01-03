package main

import "fmt"

func main_1_13_1() {
	var son int
	fmt.Scan(&son)
	fmt.Println(son/100 + son/10%10 + son%10)
}
