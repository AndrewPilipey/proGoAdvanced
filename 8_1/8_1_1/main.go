package main

import "fmt"

const (
	StatusNew       = "new"
	StatusPaid      = "paid"
	StatusShipped   = "shipped"
	StatusDelivered = "delivered"
)

func main() {
	var currentStatus string
	_, _ = fmt.Scan(&currentStatus)

	switch currentStatus {
	case StatusNew:
		fmt.Println("Заказ создан")
	case StatusPaid:
		fmt.Println("Заказ оплачен")
	case StatusShipped:
		fmt.Println("Заказ отправлен")
	case StatusDelivered:
		fmt.Println("Заказ доставлен")
	default:
		fmt.Println("Неизвестный статус")
	}

}
