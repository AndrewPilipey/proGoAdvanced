package main

import "fmt"

func print(values []int) {
	length := len(values)
	for i := 0; i < length; i++ {
		fmt.Print(values[i], " ")
	}
	fmt.Println()
}

func change(array []int) {
	lenArr := len(array)
	for i := 0; i < lenArr; i++ {
		array[i] = array[i] * 2
	}
}

func main() {
	a := []int{1, 4, 5, 7}
	print(a)
	change(a)
	print(a)
	fmt.Println("Конец main...")
}
