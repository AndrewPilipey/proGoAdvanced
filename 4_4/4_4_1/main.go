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
	line := scanner.Text()

	line = strings.ToLower(line)

	if strings.Contains(line, "хорош") {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
