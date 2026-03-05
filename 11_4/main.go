package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("1")
	}()

	panic("panic")

	defer func() {
		fmt.Println("2")
	}()
}
