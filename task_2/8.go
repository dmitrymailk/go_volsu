package main

import (
	"fmt"
	"strings"
)

func findLongestWord(text string) string {

	var longestWord string
	maxLength := 0
	// Разбиваем текст на слова
	words := strings.Fields(text)

	for _, word := range words {
		length := len(word)
		if length > maxLength {
			longestWord = word
			maxLength = length
		}
	}

	return longestWord
}

func main() {
	// 8. Реализовать процедуру поиска самого длинного слова в тексте.
	text := "Это тестовый пример в котором я ищу самое длиннннннннннное слово, еще какие-то слова."

	longest := findLongestWord(text)

	fmt.Println("Самое длинное слово:", longest)
}
