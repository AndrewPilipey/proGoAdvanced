package main

import (
	"fmt"
)

type Dish struct {
	Name  string
	Price float64
}

type Restaurant struct {
	Dishes  []Dish
	Waiters []string
}

func main() {
	salmonTartar := Dish{}
	salmonTartar.Name = "Тартар из лосося"
	salmonTartar.Price = 15.5

	beefCarpaccio := Dish{}
	beefCarpaccio.Name = "Карпаччо из говядины"
	beefCarpaccio.Price = 18.0

	beefTongue := Dish{}
	beefTongue.Name = "Язык с хреном"
	beefTongue.Price = 12.5

	winePlate := Dish{}
	winePlate.Name = "Винное плато"
	winePlate.Price = 22.0

	restaurant := Restaurant{}
	restaurant.Dishes = []Dish{salmonTartar, beefCarpaccio, beefTongue, winePlate}
	restaurant.Waiters = []string{"Валерия"}

	for _, dish := range restaurant.Dishes {
		fmt.Printf("%s — %.1f р.\n", dish.Name, dish.Price)
	}

	for _, waiter := range restaurant.Waiters {
		fmt.Printf("Официант: %s\n", waiter)
	}
}
