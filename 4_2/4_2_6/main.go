package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	for i := range s {
		if s[i] >= '0' && s[i] <= '9' {
			fmt.Print(string(s[i]), " ")
		}
	}
}
