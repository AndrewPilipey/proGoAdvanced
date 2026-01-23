package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	var result string
	for i := 0; i < len(line); i++ {
		for j := range line {
			if i != j {
				if string(line[i]) == string(line[j]) {
					result = string(line[i])
				}
			}
		}
	}
	fmt.Println(result)
}
