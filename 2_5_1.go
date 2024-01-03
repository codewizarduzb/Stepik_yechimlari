package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main_2_5_1() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	rs := []rune(text)

	if unicode.IsUpper(rs[0]) == true && string(rs[len(rs)-1]) == "." {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}

}
