package main

import (
	"fmt"
	"math"
)

// Определяем структуру "точка"
type Point struct {
	X float64
	Y float64
}

// Определяем структуру "отрезок"
type Segment struct {
	Start Point
	End   Point
}

// Метод для структуры "отрезок", возвращающий длину отрезка
func (s Segment) Length() float64 {
	return math.Sqrt(math.Pow(s.End.X-s.Start.X, 2) + math.Pow(s.End.Y-s.Start.Y, 2))
}

// Определяем структуру "треугольник"
type Triangle struct {
	A Point
	B Point
	C Point
}

// Метод для структуры "треугольник", возвращающий площадь
func (t Triangle) Area() float64 {
	return math.Abs((t.A.X*(t.B.Y-t.C.Y) + t.B.X*(t.C.Y-t.A.Y) + t.C.X*(t.A.Y-t.B.Y)) / 2.0)
}

// Определяем структуру "круг"
type Circle struct {
	Center Point
	Radius float64
}

// Метод для структуры "круг", возвращающий площадь
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Определяем интерфейс "фигура"
type Shape interface {
	Area() float64
}

// Функция для печати площади фигуры
func printArea(s Shape) {
	result := s.Area()
	fmt.Printf("Площадь фигуры: %.2f\n", result)
}

func main() {
	// Создаем экземпляры фигур
	triangle := Triangle{
		A: Point{0, 0},
		B: Point{5, 0},
		C: Point{0, 5},
	}

	circle := Circle{
		Center: Point{0, 0},
		Radius: 3,
	}

	// Вызываем функцию printArea для треугольника и круга
	printArea(triangle)
	printArea(circle)

	// Создаем отрезок и выводим его длину
	segment := Segment{
		Start: Point{0, 0},
		End:   Point{3, 4},
	}
	fmt.Printf("Длина отрезка: %.2f\n", segment.Length())
}
