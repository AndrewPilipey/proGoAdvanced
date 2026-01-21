package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	sentence := scanner.Text()

	runes := []rune(sentence)

	for i := 0; i < len(runes); i++ {
		if runes[i] == 'e' {
			runes[i] = 'i'
		}
	}
	sentence = string(runes)
	fmt.Println(sentence)
}
