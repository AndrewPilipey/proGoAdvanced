package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	received := scanner.Text()

	for i := 0; i < len(received); i++ {
		if received[i] == '.' {
			fmt.Print(0)
		} else if received[i] == '-' {
			if i+1 < len(received) {
				if received[i+1] == '.' {
					fmt.Print(1)
					i++
				} else if received[i+1] == '-' {
					fmt.Print(2)
					i++
				}

			} else {
				fmt.Println("Wrong input.")
			}
		}
	}

}
