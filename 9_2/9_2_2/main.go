package main

import (
	"fmt"
)

func main() {
	var n, m int
	_, _ = fmt.Scan(&n, &m)

	pn := &n
	pm := &m

	digitsN := countDigits(pn)
	digitsM := countDigits(pm)

	result := digitsN * digitsM
	fmt.Println(result)
}

func countDigits(p *int) int {
	if p == nil {
		return 0
	}
	temp := *p
	counter := 0
	for temp != 0 {
		temp = temp / 10
		counter++
	}
	return counter
}
