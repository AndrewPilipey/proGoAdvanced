package main

import "fmt"

type Person struct{ Name string }

func (p Person) String() string {
	return "Person: " + p.Name
}

func main() {
	p := Person{"Alex"}
	fmt.Println(p)
}
