package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()

	counterLow := 0
	counterUp := 0
	for _, char := range s {
		if unicode.IsLower(char) {
			counterLow++
		} else if unicode.IsUpper(char) {
			counterUp++
		}
	}
	if counterUp > counterLow {
		fmt.Println(strings.ToUpper(s))
	} else {
		fmt.Println(strings.ToLower(s))
	}
}
