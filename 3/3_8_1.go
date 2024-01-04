package main

// Пакет и функция main уже объявлены, выводить и считывать ничего не нужно!
func task(channel chan int, N int) chan int {
	channel <- N + 1
	return channel
}
