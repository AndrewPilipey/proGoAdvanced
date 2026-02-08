type Order struct {
	OrderedDishes []Dish  // список блюд
	TotalPrice    float64 // стоимость заказа
	OrderTime     string  // время заказа (можно заменить на time.Time)
	TableNumber   string  // номер столика
}

type Customer struct {
	FullName    string  // имя клиента
	Orders      []Order // список заказов
	PhoneNumber string  // телефон клиента
}

order := Order{[]Dish{pastaCarbonara}, 16.9, "18", "1"}
customer := Customer{"Артур", []Order{order}, "1-123-123"}




