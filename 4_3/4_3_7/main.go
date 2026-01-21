package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	sentense := scanner.Text()
	words := strings.Fields(sentense)
	result := strings.Join(words, " ")
	fmt.Println(result)
}
