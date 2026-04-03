package main

import "fmt"

func Reverse[T any](s []T) {
	for i := 0; i < len(s)/2; i++ {
		j := len(s) - 1 - i
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	Reverse(nums)
	fmt.Println(nums)
}
