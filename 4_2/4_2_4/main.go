package main

import (
	"fmt"
	"unicode"
)

func main() {
	var s string
	_, _ = fmt.Scan(&s)

	var firstRune rune
	for _, char := range s {
		firstRune = char
		break
	}

	if unicode.IsUpper(firstRune) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
