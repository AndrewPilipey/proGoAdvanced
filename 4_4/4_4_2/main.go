package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	// FIRST VERSION
	// counter := 0
	// for i := range line {
	// 	if string(line[i]) >= "a" && string(line[i]) <= "z" {
	// 		counter++
	// 	}
	// }

	// SECOND VERSION
	counter := 0
	for i := range line {
		if unicode.IsLetter(rune(line[i])) && unicode.IsLower(rune(line[i])) {
			counter++
		}
	}

	fmt.Println(counter)
}
