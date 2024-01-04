package main

// Пакет и функция main уже объявлены, выводить и считывать ничего не нужно!
func task2(channel chan string, word string) chan string {
	channel <- word + " "
	channel <- word + " "
	channel <- word + " "
	channel <- word + " "
	channel <- word + " "

	return channel
}
