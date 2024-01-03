package main

func calculator_1(arguments <-chan int, done <-chan struct{}) <-chan int {
	resultCh := make(chan int)

	go func() {
		defer close(resultCh)

		sum := 0
		doneFlag := false

		for {
			select {
			case arg := <-arguments:
				sum += arg
			case <-done:
				doneFlag = true
			}

			if doneFlag {
				resultCh <- sum
				return
			}
		}
	}()

	return resultCh
}
