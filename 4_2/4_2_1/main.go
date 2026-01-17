package main

import (
	"fmt"
	"unicode"
)

func main() {
	var a rune
	_, _ = fmt.Scanf("%c", &a)

	if unicode.IsDigit(a) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
