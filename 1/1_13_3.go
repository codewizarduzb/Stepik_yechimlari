package main

import "fmt"

func main_1_13_3() {
	var sekund int
	fmt.Scan(&sekund)

	soat := sekund / 3600
	daqiqa := sekund/60 - soat*60

	fmt.Printf("It is %d hours %d minutes.", soat, daqiqa)

}
