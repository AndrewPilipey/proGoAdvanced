package main

import (
	"cmp"
	"fmt"
)

func FindMax[T cmp.Ordered](slice []T) T {
	if len(slice) == 0 {
		var zero T
		return zero
	}
	max := slice[0]
	for _, v := range slice[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

func main() {
	nums := []int{3, 7, 2, 8, 1}
	fmt.Println(FindMax(nums))

	words := []string{"яблоко", "груша", "банан", "апельсин"}
	fmt.Println(FindMax(words))
}
