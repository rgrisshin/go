package main

import (
	"fmt"
)

// Функция Map, применяющая функцию к каждому элементу среза
func Map(input []float64, fn func(float64) float64) []float64 {
	// Создаем новый срез для результатов
	result := make([]float64, len(input))
	// Применяем функцию к каждому элементу
	for i, v := range input {
		result[i] = fn(v)
	}
	return result
}

// Функция для возведения в квадрат
func square(x float64) float64 {
	return x * x
}

func main() {
	// Создаем срез и заполняем его значениями
	values := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	fmt.Println("Исходный срез:", values)

	// Присваиваем функцию square переменной
	squareFunc := square

	// Применяем функцию Map
	squaredValues := Map(values, squareFunc)
	fmt.Println("Срез после применения функции Map:", squaredValues)
}
