package main

import (
	"fmt"
	"strings"
)

func main_2_5_3() {

	var word, something string
	fmt.Scan(&word, &something)

	fmt.Println(strings.Index(word, something))

}
