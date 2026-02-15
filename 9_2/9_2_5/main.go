package main

import "fmt"

func applyOp(op func(int, int) int, a, b int) int {
	return op(a, b)
}

func main() {
	fmt.Println("Введите операнд:")
	var operand string
	_, _ = fmt.Scan(&operand)
	fmt.Println("Введите два числа:")
	var a, b int
	_, _ = fmt.Scan(&a, &b)

	var operation func(int, int) int
	switch operand {

	case "+":
		operation = func(x, y int) int { return x + y }
	case "-":
		operation = func(x, y int) int { return x - y }
	case "*":
		operation = func(x, y int) int { return x * y }
	case "%":
		operation = func(x, y int) int { return x % y }
	}

	result := applyOp(operation, a, b)
	fmt.Println(result)
}
