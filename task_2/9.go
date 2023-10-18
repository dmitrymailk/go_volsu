package main

import (
	"fmt"
	"strings"
)

func countWordsWithLength(text string, length int) int {

	var count int

	words := strings.Fields(text)

	for _, word := range words {
		fmt.Printf("%s: %d\n", word, len(word))
		if len(word) == length {
			count++
		}
	}

	return count
}

func main() {
	// 9. Реализовать процедуру подсчета числа слов в тексте с заданным количеством букв.
	text := "This is my test sentence for example for task 9."
	length := 3
	count := countWordsWithLength(text, length)

	fmt.Printf("Количество слов длины %d: %d\n", length, count)
}
