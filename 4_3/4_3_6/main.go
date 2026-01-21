package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	nStr := scanner.Text()

	n, _ := strconv.Atoi(nStr)
	slice := make([]string, n)

	for i := range slice {
		scanner.Scan()
		slice[i] = scanner.Text()
	}

	prefix := slice[0]

	for i := 1; i < len(slice); i++ {
		for !strings.HasPrefix(slice[i], prefix) {
			prefix = prefix[:len(prefix)-1]
			if len(prefix) == 0 {
				fmt.Println("НЕТ")
				break
			}
		}
	}
	fmt.Println(prefix)
}
