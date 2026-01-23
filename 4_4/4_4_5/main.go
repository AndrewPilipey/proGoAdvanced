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
	line1 := scanner.Text()
	scanner.Scan()
	line2 := scanner.Text()

	line1, line2 = strings.ToLower(line1), strings.ToLower(line2)

	if line1 < line2 {
		fmt.Println(-1)
	} else if line1 > line2 {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
