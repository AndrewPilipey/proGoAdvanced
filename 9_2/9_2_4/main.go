package main

import (
	"fmt"
)

func main() {
	a, b := 1, 2
	swapValues(&a, &b)
	fmt.Println(a, b)
}

func swapValues(a *int, b *int) {
	*a, *b = *b, *a
}
