package main

import (
	"fmt"
)

type Student struct {
	Name   []string
	Number []int
}

func Find[T comparable](slice []T, target T) int { //to search target accordingly
	for i, v := range slice {
		if v == target {
			return i
		}
	}
	return -1
}

func main() {
	var searchMode string
	_, _ = fmt.Scan(&searchMode) //searchMode: name or number or student

	var n int
	_, _ = fmt.Scan(&n) // n: the number of students

	var foundedIdx int

	students := Student{
		Name:   make([]string, n),
		Number: make([]int, n),
	}

	//var nReduserer int = n
	for i := 0; i < n; i++ {
		//	fmt.Printf("You have to enter n-names (%d/%d) of students: ", nReduserer, n)
		_, _ = fmt.Scan(&students.Name[i], &students.Number[i])
		//	nReduserer -= 1

	}

	//fmt.Println("Target: ")
	switch searchMode {
	case "name":
		var targetName string
		_, _ = fmt.Scan(&targetName)
		foundedIdx = Find(students.Name, targetName)

	case "number":
		var targetNum int
		_, _ = fmt.Scan(&targetNum)
		foundedIdx = Find(students.Number, targetNum)

	case "student":
		var mark bool = true
		var name string
		_, _ = fmt.Scan(&name)
		foundedIdx = Find(students.Name, name)
		if foundedIdx == -1 {
			mark = false
		}

		var num int
		_, _ = fmt.Scan(&num)
		foundedIdx = Find(students.Number, num)
		if foundedIdx == -1 {
			mark = false
		}
		if mark == true && foundedIdx != -1 {
			fmt.Println(foundedIdx)
			return // выход
		}
	default:
		fmt.Println("Unknown type:", searchMode)
		return
	}

	if foundedIdx != -1 {
		fmt.Println(foundedIdx)
	} else {
		fmt.Println(-1)
	}
}
