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

	globalCounter := 0
	counter := 0

	for i := range line {
		if string(line[i]) == "x" {
			counter++
		} else if string(line[i]) != "x" {
			counter = 0
		}
		if counter >= 3 {
			globalCounter += 1
		}
	}

	fmt.Println(globalCounter)

}
