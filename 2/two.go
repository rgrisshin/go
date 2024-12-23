package main

import (
	"fmt"
)

// Функция countCharacters принимает строку и возвращает карту с количеством вхождений каждого символа
func countCharacters(s string) map[rune]int {
	charCount := make(map[rune]int)

	for _, char := range s {
		charCount[char]++
	}

	return charCount
}

func main() {
	// Пример строки
	input := "Привет, мир!"
	// Вызов функции и получение результата
	charactersCount := countCharacters(input)

	// Вывод результата
	for char, count := range charactersCount {
		fmt.Printf("%c: %d\n", char, count)
	}
}
