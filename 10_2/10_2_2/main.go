package main

import "fmt"

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}

type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

func main() {
	circle := Circle{Radius: 3.0}
	rectangle := Rectangle{Width: 4.0, Height: 5.0}
	triangle := Triangle{Base: 6.0, Height: 4.0}
	square := Square{Side: 4.0}

	shapes := []Shape{circle, rectangle, triangle, square}

	fmt.Println("Площади фигур:")
	for _, shape := range shapes {
		fmt.Printf("%.2f\n", shape.Area())
	}
}
