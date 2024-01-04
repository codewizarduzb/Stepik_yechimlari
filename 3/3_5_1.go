package main

import (
	"bufio"
	"os"
	"strconv"
)

// package main уже объявлен.
func main_3_5_1() {
	// Все указанные в тексте задачи пакеты уже импортированы (strconv, bufio, os, io)
	// Другие использовать нельзя.
	s := bufio.NewScanner(os.Stdin)

	var sum int

	for s.Scan() {
		line := s.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			os.Stderr.WriteString("Error")
			continue
		}
		sum += num
	}

	sumStr := strconv.Itoa(sum)
	os.Stdout.WriteString(sumStr)
}
