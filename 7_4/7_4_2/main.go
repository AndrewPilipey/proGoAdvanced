package main

type Dish struct {
	Name     string
	Calories int
}

type Restaurant struct {
	Address string
	Menu    []Dish
}

func main() {
	restaurant := Restaurant{
		Address: "ул. Примерная, 10",
		Menu:    []Dish{},
	}

	pastaCarbonara := Dish{"Паста карбонара", 550}
	restaurant.Menu = append(restaurant.Menu, pastaCarbonara)

	meatballSoup := Dish{"Суп с фрикадельками", 300}
	restaurant.Menu = append(restaurant.Menu, meatballSoup)
}
