package main

import (
	"fmt"
	"unicode"
)

func main() {
	var c rune
	_, _ = fmt.Scanf("%c", &c)

	if unicode.IsUpper(c) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
