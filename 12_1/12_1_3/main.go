package main

import "fmt"

func Swap[T any](a, b T) (T, T) {
	return b, a
}

func main() {
	x, y := 1, 2
	x, y = Swap(x, y)
	fmt.Println(x, y)

	s1, s2 := "hello", "world"
	s1, s2 = Swap(s1, s2)
	fmt.Println(s1, s2)

	a, b := 3.14, 2.71
	a, b = Swap(a, b)
	fmt.Println(a, b)

	d, f := 10, 20
	d, f = Swap(d, f)
	fmt.Println(d, f)
}
