package main

import "fmt"

func main() {
	var s string
	_, _ = fmt.Scan(&s)

	counterPluses := 0
	counterStars := 0

	for i := range s {
		if s[i] == '+' {
			counterPluses++
		}
		if s[i] == '*' {
			counterStars++
		}
	}

	fmt.Printf("Символ + встречается %d раз\n", counterPluses)
	fmt.Printf("Символ * встречается %d раз", counterStars)

}
