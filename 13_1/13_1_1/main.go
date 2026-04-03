package main

import (
	"fmt"
	"strconv"
)

func FindInt(slice []int, value int) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}

func FindString(slice []string, value string) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}

func main() {
	var n int
	fmt.Scan(&n)

	var firstEl string
	fmt.Scan(&firstEl)

	_, err := strconv.Atoi(firstEl)
	isNumber := err == nil

	if isNumber {

		intSlice := make([]int, n)
		firstNum, _ := strconv.Atoi(firstEl)
		intSlice[0] = firstNum

		for i := 1; i < n; i++ {
			var numStr string
			fmt.Scan(&numStr)
			num, _ := strconv.Atoi(numStr)
			intSlice[i] = num
		}

		var targetStr string
		fmt.Scan(&targetStr)
		targetNum, _ := strconv.Atoi(targetStr)

		result := FindInt(intSlice, targetNum)
		fmt.Println(result)
	} else {

		stringSlice := make([]string, n)
		stringSlice[0] = firstEl

		for i := 1; i < n; i++ {
			fmt.Scan(&stringSlice[i])
		}

		var targetString string
		fmt.Scan(&targetString)

		result := FindString(stringSlice, targetString)
		fmt.Println(result)
	}
}
