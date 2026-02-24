package main

import "fmt"

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Triangle struct {
	Base   float64
	Height float64
}

type Square struct {
	Side float64
}

func main() {
	circle := Circle{Radius: 3.0}
	rectangle := Rectangle{Width: 4.0, Height: 5.0}
	triangle := Triangle{Base: 6.0, Height: 4.0}
	square := Square{Side: 4.0}

	fmt.Printf("Circle: Radius = %.2f\n", circle.Radius)
	fmt.Printf("Rectangle: Width = %.2f, Height = %.2f\n", rectangle.Width, rectangle.Height)
	fmt.Printf("Triangle: Base = %.2f, Height = %.2f\n", triangle.Base, triangle.Height)
	fmt.Printf("Square: Side = %.2f\n", square.Side)
}
