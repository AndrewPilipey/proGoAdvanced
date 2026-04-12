package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
	"unicode/utf8"
)

func getLastRune(s2 string) (rune, bool) {
	if len(s2) == 0 {
		return 0, false
	}
	r, _ := utf8.DecodeLastRuneInString(s2)
	return r, true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	s1 := scanner.Text()

	scanner.Scan()
	s2 := scanner.Text()

	if len(s1) == 0 || len(s2) == 0 {
		//	fmt.Println("Error: empty string")
		return
	}

	s1First := []rune(s1)[0]
	s2last, ok := getLastRune(s2)

	if !ok {
		fmt.Println("NO")
		return
	}

	if unicode.ToLower(s1First) == unicode.ToLower(s2last) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
