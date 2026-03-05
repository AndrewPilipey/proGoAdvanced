// Напишите функцию swap(a, b *int),
// которая меняет местами значения двух переменных,
// используя указатели.

package main

import "fmt"

func swapNums(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	var a, b int
	_, _ = fmt.Scan(&a, &b)

	fmt.Printf("До: a = %d, b = %d\n", a, b) //5 10

	swapNums(&a, &b)
	fmt.Printf("После: a = %d, b = %d", a, b) // 10 5
}
