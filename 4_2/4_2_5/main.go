package main

import (
	"fmt"
	"unicode"
)

func main() {
	var s1, s2 string
	_, _ = fmt.Scan(&s1, &s2)

	firstRune := getFirstRune(s1)
	lastRune := getLastRune(s2)

	unicode.ToLower(firstRune)
	unicode.ToLower(lastRune)

	if string(firstRune) == string(lastRune) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func getFirstRune(s string) rune {
	for _, char := range s {
		return char
	}
	return -1
}

func getLastRune(s string) rune {
	var last rune
	for _, char := range s {
		last = char
	}
	return last
}
