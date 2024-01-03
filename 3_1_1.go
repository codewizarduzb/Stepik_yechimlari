package main

import "fmt"

func main_3_1_1() {
	var son int
	m := make(map[int]int)

	for i := 0; i < 10; i++ {
		fmt.Scan(&son)
		if key, ok := m[son]; ok {
			fmt.Print(key, " ")
		} else {
			//		m[son] = work(son)

			fmt.Print(m[son], " ")
		}
	}
}
