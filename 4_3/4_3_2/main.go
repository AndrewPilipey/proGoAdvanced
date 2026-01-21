package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()

	list := strings.Split(s, " ")
	name := list[0]
	surname := list[1]

	if unicode.IsUpper(rune(name[0])) && unicode.IsUpper(rune(surname[0])) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
