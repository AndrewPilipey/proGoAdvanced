package main

import (
	"cmp"
	"fmt"
)

func main() {
	var DataType string
	//fmt.Println("Enter a datatype: int, float OR string")
	_, _ = fmt.Scan(&DataType)

	var n int
	//fmt.Println("Enter the number of elements:")
	_, _ = fmt.Scan(&n)

	switch DataType {
	case "int":
		slice := make([]int, n)
		slice = Scan(slice)
		result := IsSorted(slice)
		if result {
			fmt.Println("YES")
			return
		} else {
			fmt.Println("NO")
			return
		}
	case "float":
		slice := make([]float64, n)
		slice = Scan(slice)
		result := IsSorted(slice)
		if result {
			fmt.Println("YES")
			return
		} else {
			fmt.Println("NO")
			return
		}
	case "string":
		slice := make([]string, n)
		slice = Scan(slice)
		result := IsSorted(slice)
		if result {
			fmt.Println("YES")
			return
		} else {
			fmt.Println("NO")
			return
		}
	}
}

func Scan[T any](slice []T) []T {
	for i := 0; i < len(slice); i++ {
		_, _ = fmt.Scan(&slice[i])
	}
	return slice
}

func IsSorted[T cmp.Ordered](slice []T) bool {
	if len(slice) == 0 || len(slice) == 1 {
		return true
	}
	for i := 0; i < len(slice)-1; i++ {
		if slice[i] > slice[i+1] {
			return false
		}
	}

	return true
}
