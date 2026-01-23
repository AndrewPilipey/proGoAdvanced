package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nStr := scanner.Text()
	n, err := strconv.Atoi(nStr)
	if err != nil {
		fmt.Println("Input error. Err:", err)
		return
	}

	var inputLines []string
	for i := 0; i < n; i++ {
		scanner.Scan()
		inputLines = append(inputLines, scanner.Text())
	}

	var results []string
	for _, line := range inputLines {
		var digits []rune
		for _, ch := range line {
			if unicode.IsDigit(ch) {
				digits = append(digits, ch)
			}
		}
		formatted := fmt.Sprintf("%c-%c%c%c-%c%c%c-%c%c-%c%c",
			digits[0], digits[1], digits[2], digits[3],
			digits[4], digits[5], digits[6],
			digits[7], digits[8],
			digits[9], digits[10])

		results = append(results, formatted)
	}

	for _, res := range results {
		fmt.Println(res)
	}
}
