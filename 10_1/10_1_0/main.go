package main

import "fmt"

// Интерфейс Shape задаёт контракт: любая фигура должна уметь вычислять площадь
type Shape interface {
	Area() float64
}

// Структура для круга
type Circle struct {
	Radius float64
}

// Реализация метода Area для круга
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Структура для прямоугольника
type Rectangle struct {
	Width, Height float64
}

// Реализация метода Area для прямоугольника
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	// Создаём фигуры
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 4, Height: 6}

	// Храним их в одном списке как интерфейсы Shape
	shapes := []Shape{circle, rectangle}

	// Единообразно вычисляем площади
	for _, shape := range shapes {
		fmt.Printf("Площадь: %.2f\n", shape.Area())
	}
}
