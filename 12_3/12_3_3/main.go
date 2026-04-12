package main

import (
	"fmt"
)

func main() {
	var DataType string
	//fmt.Println("Enter DataType: int / float64 / string")
	_, _ = fmt.Scan(&DataType)

	var n int
	//fmt.Println("Enter the number of elements:")
	_, _ = fmt.Scan(&n)

	switch DataType {
	case "int":
		slice := make([]int, n)
		slice = Scan(slice)
		result := Unique(slice)
		for i, v := range result { // print result
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(v)
		}
		fmt.Println()
	case "string":
		slice := make([]string, n)
		slice = Scan(slice)
		result := Unique(slice)
		for i, v := range result { // print result
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(v)
		}
	case "float":
		slice := make([]float64, n)
		slice = Scan(slice)
		result := Unique(slice)
		for i, v := range result { // print result
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(v)
		}
	default:
		fmt.Println("The wrong type of data")
		return
	}

}

func Scan[T any](slice []T) []T {
	for i := 0; i < len(slice); i++ {
		_, _ = fmt.Scan(&slice[i])
	}
	return slice
}

func Unique[T comparable](slice []T) []T {
	seen := make(map[T]bool)
	result := []T{}
	for _, v := range slice {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}
