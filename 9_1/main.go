package main

import (
	"fmt"
)

func main() {
	var paper *int
	fmt.Println(paper)

	box := 42
	paper = &box
	fmt.Println(paper, *paper, box)

	whatsInside := *paper
	fmt.Println(whatsInside)
	//whatsInside = *paper
	fmt.Println("---------")
	var p *int
	fmt.Println(p, p)
	p = &box
	fmt.Println(p, *p)
	fmt.Println("---------")
	p2 := new(int)
	*p2 = 100 // безопасно, память есть
	fmt.Println(p, *p, p2, *p2, box)
	p2 = p
	fmt.Println(p, *p, p2, *p2, box)

}
