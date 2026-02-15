package main

import (
	"fmt"
)

func main() {
	x, y := 5, 8
	setToMax(&x, &y)
	fmt.Println(x)
}

func setToMax(a *int, b *int) {
	if a == nil || b == nil {
		return
	}

	if *a < *b {
		*a = *b
	}
}
