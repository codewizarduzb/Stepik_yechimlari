package main

import "fmt"

func main_2_5_2() {

	var word, new_word string

	fmt.Scan(&word)

	rs := []rune(word)

	for i := len(rs) - 1; i >= 0; i-- {
		new_word = new_word + string(rs[i])
	}

	if word == new_word {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}
}
