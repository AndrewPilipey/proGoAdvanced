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
	fmt.Println(s)

	var pointDeleting int

	_, _ = fmt.Scan(&pointDeleting)

	// if err != nil {

	//  fmt.Println("Wrong input")

	//  return

	// }

	// if pointDeleting < 0 {

	//  fmt.Println("negative index")

	// }

	if pointDeleting == len(s) {

		pointDeleting = len(s) - 1

	}

	s = s[:pointDeleting] + s[pointDeleting+1:]

	fmt.Println(s)

}
