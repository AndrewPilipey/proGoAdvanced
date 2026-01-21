package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var lyrics string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		lyrics = scanner.Text()
	} else {
		fmt.Println("wrong input")
	}

	lyrics = strings.ReplaceAll(lyrics, "WUB", " ")
	result := strings.Fields(lyrics)
	fmt.Println(strings.Join(result, " "))

}
