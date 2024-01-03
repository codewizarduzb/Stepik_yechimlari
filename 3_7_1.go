package main

import (
	"fmt"
	"time"
)

func main_3_7_1() {

	var inputDate string
	fmt.Scan(&inputDate)

	parsedTime, err := time.Parse(time.RFC3339, inputDate)
	if err != nil {
		panic(err)
	}
	fmt.Println(parsedTime.Format(time.UnixDate))

}
