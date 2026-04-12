package main

import (
	"fmt"
)

func main() {
	var num int
	_, _ = fmt.Scan(&num)
	Identity(num)

	var str string
	_, _ = fmt.Scan(&str)
	Identity(str)

	var ship float64
	_, _ = fmt.Scan(&ship)
	Identity(ship)

}

func Identity[T any](value T) {
	fmt.Printf("%v %T\n", value, value)
}
