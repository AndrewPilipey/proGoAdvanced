package main

import (
	"fmt"
)

type Table struct {
	Number int
	Guests []string
}

type Waiter struct {
	Name   string
	Tables []Table
}

func main() {
	tableA := Table{}
	tableA.Number = 4
	tableA.Guests = []string{"Гость 4"}

	tableB := Table{}
	tableB.Number = 7
	tableB.Guests = []string{"Гость 12"}

	valeria := Waiter{}
	valeria.Name = "Валерия"
	valeria.Tables = []Table{tableA}

	konstantin := Waiter{}
	konstantin.Name = "Константин"
	konstantin.Tables = []Table{tableB}

	fmt.Printf("Официант %s обслуживает стол %d\n", valeria.Name, valeria.Tables[0].Number)
	fmt.Printf("Официант %s обслуживает стол %d\n", konstantin.Name, konstantin.Tables[0].Number)
}
