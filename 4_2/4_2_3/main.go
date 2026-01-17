package main

import (
	"fmt"
	"unicode"
)

func main() {
	var c rune
	_, _ = fmt.Scanf("%c", &c)

	if unicode.IsLetter(c) {
		if unicode.IsLower(c) {
			c = unicode.ToUpper(c)
		} else {
			c = unicode.ToLower(c)
		}
	}

	fmt.Println(string(c))
}
