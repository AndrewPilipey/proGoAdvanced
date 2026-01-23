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

	resultX := strings.Index(line, "x")
	resultW := strings.Index(line, "w")

	if resultX == -1 && resultW == -1 {
		fmt.Println(-1)
		return
	}

	if resultX == -1 {
		resultX = len(line)
	}
	if resultW == -1 {
		resultW = len(line)
	}

	result := min(resultX, resultW)

	fmt.Println(string(line[result]))

}
