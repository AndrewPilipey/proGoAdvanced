package main

import "fmt"

func update(s *string) {
	if s == nil {
		fmt.Println("nil")
		return
	}
	*s = *s + "!"
}

func main() {
	var text *string
	update(text)

	msg := "Hello"
	update(&msg)
	fmt.Println(msg)
}
