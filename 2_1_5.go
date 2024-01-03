package main

func sumInt(args ...int) (int, int) {
	sum := 0
	count := 0
	for _, value := range args {
		sum += value
		count++
	}
	return count, sum
}
