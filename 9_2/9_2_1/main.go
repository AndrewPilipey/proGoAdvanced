package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	scanner.Scan()
	symbol := scanner.Text()[0]

	ptr := &input

	result := isCharacterUnique(ptr, symbol)
	fmt.Println(result)
}

func isCharacterUnique(p *string, s byte) bool {
	if p == nil {
		return false
	}
	unique := 0
	point := *p
	for i := range point {
		if point[i] == s { // можно ещё без str, через (*p)[i]
			unique++ // однако (*p)[i] затратнее по времени,
			// поскольку разыменование происходит
			// каждый раз.
		}
	}
	if unique == 1 {
		return true
	}
	return false
}
