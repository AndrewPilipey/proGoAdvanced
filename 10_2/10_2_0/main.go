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
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Triangle struct {
	Base, Height float64
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height / 2
}

// /*
func PrintArea(shape Shape) {
	fmt.Printf("Площадь фигуры: %.2f\n", shape.Area())
}

func TotalArea(shapes []Shape) float64 {
	var total float64
	for _, shape := range shapes {
		total += shape.Area()
	}
	return total
}

// */
func main() {
	var shapes []Shape = []Shape{
		Circle{Radius: 5},
		Rectangle{Width: 4, Height: 6},
		Triangle{Base: 3, Height: 4},
	}

	for _, shape := range shapes {
		PrintArea(shape)
	}
	fmt.Printf("Площадь всех фигур: %.2f\n", TotalArea(shapes))

}
