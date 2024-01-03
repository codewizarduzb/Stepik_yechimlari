package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main_3_7_3() {
	var sana = []string{}
	input := bufio.NewScanner(os.Stdin)

	if input.Scan() {
		sana = strings.Split(input.Text(), ",")
	}

	first_date := sana[0]
	second_date := sana[1]

	date_1, err_1 := time.Parse("02.01.2006 15:04:05", first_date)
	if err_1 != nil {
		panic(err_1)
	}
	date_2, err_2 := time.Parse("02.01.2006 15:04:05", second_date)
	if err_2 != nil {
		panic(err_2)
	}

	if date_1.Before(date_2) {
		fmt.Println(date_2.Sub(date_1))
	} else {
		fmt.Println(date_1.Sub(date_2))
	}
}
