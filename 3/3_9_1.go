package main

import "time"

func main_3_9_1() {
	// вы уже внутри main()
	go func() {
		//		work()
	}()

	time.Sleep(time.Second)
}
