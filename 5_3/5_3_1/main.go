package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var sep string
	scanner.Scan()
	sep = scanner.Text()

	news := [3]string{}

	for i := range news {
		scanner.Scan()
		news[i] = scanner.Text()
	}

	result := fmt.Sprintf("%s%s%s%s%s", news[0], sep, news[1], sep, news[2])
	fmt.Println(result)
}
