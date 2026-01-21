//в вскод редакторе всё рабоатет норм при любом из
// подходов, а степик так и не принял

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()

	var pointDel int
	_, _ = fmt.Scan(&pointDel)

	runes := []rune(s)
	result := string(runes[:(pointDel-1)]) + string(runes[pointDel:])

	fmt.Println(result)
	// for i := 0; i < len(s); i++ {
	// 	if i == pointDel-1 {
	// 		continue
	// 	} else {
	// 		fmt.Print(string(s[i]))
	// 	}
	// }
}
