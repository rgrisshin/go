package main

import (
	"errors"
	"fmt"
)

// Функция formatIP принимает массив из 4 байтов и возвращает строку с IP-адресом
func formatIP(ip [4]byte) string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

// Функция listEven принимает диапазон и возвращает срез четных чисел и ошибку
func listEven(start, end int) ([]int, error) {
	if start > end {
		return nil, errors.New("левая граница больше правой")
	}

	var evens []int
	for i := start; i <= end; i++ {
		if i%2 == 0 {
			evens = append(evens, i)
		}
	}
	return evens, nil
}

func main() {
	// Пример использования функции formatIP
	ip := [4]byte{127, 0, 0, 1}
	fmt.Println("IP-адрес:", formatIP(ip))

	// Пример использования функции listEven
	evens, err := listEven(1, 10)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Чётные числа:", evens)
	}

	// Пример с ошибкой
	evens, err = listEven(10, 1)
	if err != nil {
		fmt.Println("Ошибка:", err)
	}
}
