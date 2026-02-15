package main

import "fmt"

type Zebra struct {
	Counter int
}

func (z *Zebra) GetStripe() {

	if z.Counter%2 == 0 {
		fmt.Println("Полоска белая")
	} else {
		fmt.Println("Полоска черная")
	}
	z.Counter++
}

func main() {
	zebra1 := Zebra{}
	zebra1.GetStripe()
	zebra1.GetStripe()
	zebra1.GetStripe()

	fmt.Println()

	zebra2 := Zebra{}
	zebra2.GetStripe()
	zebra2.GetStripe()
}
