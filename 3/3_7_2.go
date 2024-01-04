package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main_3_7_2() {

	input := bufio.NewScanner(os.Stdin)

	if input.Scan() {
		date := input.Text()

		salom, err := time.Parse("2006-01-02 15:04:05", date)
		if err != nil {
			log.Fatal(err)
		}
		if salom.Hour() < 13 || salom.Hour() == 13 && salom.Minute() == 0 {
			fmt.Println(salom.Format("2006-01-02 15:04:05"))
		} else {
			fmt.Println(salom.AddDate(0, 0, 1).Format("2006-01-02 15:04:05"))
		}
	}

}
