package main

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	resultChan := make(chan int, 1)

	go func() {
		defer close(resultChan)
		select {
		case value_1 := <-firstChan:
			resultChan <- (value_1 * value_1)
		case value_2 := <-secondChan:
			resultChan <- (value_2 * 3)
		case <-stopChan:
			break
		}
	}()

	return resultChan
}
