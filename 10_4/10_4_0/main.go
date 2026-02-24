package main

import "fmt"

func main() {
	var data any = 100

	slice := data.(float64)

	fmt.Println(slice)
}
