package main

import "fmt"

func mysteriousIncr(val *int) {
	temp := val
	*temp++

}

func main() {
	var x int
	_, _ = fmt.Scan(&x)
	fmt.Println("До:", x)

	mysteriousIncr(&x)

	fmt.Println("После:", x)
}
