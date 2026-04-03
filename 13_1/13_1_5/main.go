package main

import (
	"fmt"
)

func First[T any](slice []T) T {
	return slice[0]
}

func main() {
	var dataType string
	_, _ = fmt.Scan(&dataType)
	var n int
	_, _ = fmt.Scan(&n)

	var slice any

	switch dataType {
	case "int":
		intSlice := make([]int, n)
		for i := range intSlice {
			_, _ = fmt.Scan(&intSlice[i])
		}
		slice = intSlice
	case "string":
		strSlice := make([]string, n)
		for i := range strSlice {
			_, _ = fmt.Scan(&strSlice[i])
		}
		slice = strSlice
	}

	var result any
	switch v := slice.(type) {
	case []int:
		result = First(v)
	case []string:
		result = First(v)
	default:
		return
	}

	fmt.Println(result)

}
