package main

import (
	"fmt"
)

func getData[T any](slice []T) {
	for i := range slice {
		_, _ = fmt.Scan(&slice[i])
	}
}

func Find[T comparable](slice []T, target T) int {
	for i, v := range slice {
		if v == target {
			return i
		}
	}
	return -1
}

func main() {
	var t string
	//	fmt.Println("Enter 'name' or 'number' for starting search:")
	_, _ = fmt.Scan(&t)

	var n int
	//	fmt.Println("Enter quantity of students:")
	_, _ = fmt.Scan(&n)

	var foundedIdx int

	switch t {
	case "name":
		//	fmt.Printf("You have to enter names of &d student\n", n)
		names := make([]string, n)
		getData(names)

		var targetName string
		_, _ = fmt.Scan(&targetName)
		foundedIdx = Find(names, targetName)

	case "number":
		//	fmt.Printf("You have to enter  IDs of %d students\n", n)
		numbers := make([]int, n)
		getData(numbers)

		var targetNumber int
		_, _ = fmt.Scan(&targetNumber)
		foundedIdx = Find(numbers, targetNumber)

	default:
		fmt.Println("Unknown type:", t)
		return
	}

	if foundedIdx != -1 {
		fmt.Println(foundedIdx)
	} else {
		fmt.Println(-1)
	}
}
