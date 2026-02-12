package main

import (
	"fmt"
)

type Product struct {
	Name  string
	Price float64
}

type Store struct {
	Products []Product
	Basket   []Product
	Orders   []Order
}

type Order struct {
	Products  []Product
	FullPrice float64
}

type Customer struct {
	Name           string
	PreviousOrders []Order
	Store          *Store
}

func (p Product) Print() {
	fmt.Printf("%s %.2f\n", p.Name, p.Price)
}

func NewStore() Store {
	return Store{
		Products: []Product{
			{"Хлеб", 25},
			{"Молоко", 100},
			{"Печенье", 50},
			{"Масло", 250},
			{"Йогурт", 300},
			{"Сок", 80},
		},
		Basket: []Product{},
	}
}

func (s *Store) AddNewProduct() {
	var nameProduct string
	var costProduct float64

point00:
	fmt.Println("Введите название продукта:")
	_, err := fmt.Scanln(&nameProduct)
	if err != nil {
		fmt.Println("Ошибка ввода названия продукта. Повторите ввод.")
		goto point00
	}

point01:
	fmt.Println("Введите стоимость продукта:")
	_, err = fmt.Scanln(&costProduct)
	if err != nil || costProduct <= 0 {
		fmt.Println("Ошибка ввода стоимости продукта. Повторите ввод.")
		goto point01
	}
	newProduct := Product{
		Name:  nameProduct,
		Price: costProduct,
	}
	s.Products = append(s.Products, newProduct)
	fmt.Printf("Продукт %s стоимостью %.2f руб. добавлен в каталог.\n",
		nameProduct, costProduct)

}

func (s Store) ShowCatalog() {
	fmt.Println("Каталог продуктов:")
	s.ShowProducts(s.Products)
}

func (s Store) ShowBasket() {

	fmt.Println("Содержимое корзины:")
	if len(s.Basket) == 0 {
		fmt.Println("Ваша корзина пока пуста.")
	}
	s.ShowProducts(s.Basket)
}

func (s Store) ShowProducts(products []Product) {
	number := 1
	for _, product := range products {
		fmt.Printf("%d. ", number)
		product.Print()
		number++
	}
}

func (s Store) AddToBasket(productNumber int) Store {
	if productNumber > len(s.Products) {
		fmt.Println("Неверный выбор продукта. Повторите ввод.")
		return s
	}
	s.Basket = append(s.Basket, s.Products[productNumber-1])
	fmt.Printf("Продукт %s успешно добавлен в корзину.\n",
		s.Products[productNumber-1].Name)
	fmt.Printf("В корзине %d продуктов.\n", len(s.Basket))
	return s
}

func NewOrder(products []Product) Order {
	var sum float64
	for _, p := range products {
		sum += p.Price
	}
	return Order{Products: products, FullPrice: sum}
}

func (s *Store) CreateOrder() {
	if len(s.Basket) == 0 {
		fmt.Println("Вы не добавили продукты в корзину. Начнём сначала.")
		return
	}
	order := NewOrder(s.Basket)
	s.Orders = append(s.Orders, order)
	s.Basket = []Product{}
	fmt.Printf("Заказ оформлен. Итоговая сумма: %.2f\n", order.FullPrice)
}

func (c *Customer) SaveOrder(store *Store) {
	if len(store.Orders) == 0 {
		fmt.Println("Нет оформленных заказов для сохранения")
		return
	}
	if len(c.PreviousOrders) == 0 {
		fmt.Println("Нет ранее оформленных заказов.")
	}
	lastOrder := store.Orders[len(store.Orders)-1]
	c.PreviousOrders = append(c.PreviousOrders, lastOrder)
	fmt.Printf("Заказ на сумму %.2f сохранён в историю.\n", lastOrder.FullPrice)
	store.Orders = []Order{}
	store.Basket = []Product{}
	fmt.Println("Заказ доставлен. Можете оформить новый заказ.")
}

func (c Customer) ShowHistory() {
	if len(c.PreviousOrders) == 0 {
		fmt.Println("У Вас нет сохранённых заказов")
		return
	}

	fmt.Println("История Ваших заказов:")
	for i, order := range c.PreviousOrders {
		fmt.Printf("Заказ #%d. Сумма: %.2f руб.\n", i+1, order.FullPrice)
		fmt.Println("Состав заказа:")
		for _, product := range order.Products {
			fmt.Printf("%-5s: %.2f руб.\n", product.Name, product.Price)
		}
		fmt.Println("---")
	}
}

func main() {
	onlineStore := NewStore()
	var productNumber int
	var answer int

	customer := Customer{
		Name:           "Игнатий",
		PreviousOrders: []Order{},
	}

	fmt.Println("Здравствуйте. Выберите действие:")
point0:
	actionNumber := showNavigationMenu()

	switch actionNumber {
	case 1:
	point1:
		onlineStore.ShowCatalog()
		fmt.Println("Выберите номер продукта.")
		_, _ = fmt.Scan(&productNumber)
		onlineStore = onlineStore.AddToBasket(productNumber)
		fmt.Printf("%-5s %-5s", "1.Продолжить выбор", "Any.Вернуться в меню.\n")
		_, _ = fmt.Scan(&answer)
		if answer == 1 {
			goto point1
		} else {
			goto point0
		}
	case 2:
		onlineStore.ShowBasket()
		fmt.Printf("%-5s", "Any.Вернуться в меню.\n")
		_, _ = fmt.Scan(&answer)
		goto point0

	case 3:
		onlineStore.CreateOrder()
		customer.SaveOrder(&onlineStore) //передаю указатель
		fmt.Printf("%-5s", "Any.Вернуться в меню.\n")
		_, _ = fmt.Scan(&answer)
		goto point0
	case 4:
		customer.ShowHistory()
		fmt.Printf("%-5s", "Any.Вернуться в меню.\n")
		_, _ = fmt.Scan(&answer)
		goto point0
	case 9:
		fmt.Println("Введите логин и пароль.")
		var login, password string

		fmt.Println("Логин: ")
		_, _ = fmt.Scan(&login)

		fmt.Println("Пароль: ")
		_, _ = fmt.Scan(&password)

		if login == "admin" || password == "admin" {
			onlineStore.AddNewProduct()
			fmt.Printf("%-5s", "Any.Вернуться в меню.\n")
			_, _ = fmt.Scan(&answer)
			goto point0
		} else {
			fmt.Println("Неверные логин и пароль. Завершение программы.")
			return
		}
	case 0:
		return
	}

}

func showNavigationMenu() int {
	fmt.Println("1. Добавить продукты в корзину.")
	fmt.Println("2. Показать корзину.")
	fmt.Println("3. Оформить заказ.")
	fmt.Println("4. Посмотреть предыдущие заказы.")
	fmt.Println("9. Доступ для менеджеров.")
	fmt.Println("0. Выход")
	var actionNumber int
	_, _ = fmt.Scan(&actionNumber)
	return actionNumber

}
