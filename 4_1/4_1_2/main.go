package main

import "fmt"

func main() {
	var s string
	_, _ = fmt.Scan(&s)

	sameNumCounter := 0

	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			sameNumCounter++
		}
	}

	fmt.Println(sameNumCounter)

}
