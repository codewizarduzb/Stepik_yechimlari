package main

// Пакет и функция main уже объявлены.
// Выводить или вводить ничего не нужно!
func removeDuplicates(inputStream <-chan string, outputStream chan<- string) {
	var prev string
	for value := range inputStream {
		if value != prev {
			outputStream <- value
			prev = value
		}
	}
	close(outputStream)
}
