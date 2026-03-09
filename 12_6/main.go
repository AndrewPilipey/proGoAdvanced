package main

import (
	"errors"
	"fmt"
)

var ErrOrderNotFound = errors.New("order not found")

type OutOfStockError struct {
	Product   string
	Available int
	Requested int
}

func (e OutOfStockError) Error() string {
	return fmt.Sprintf("out of stock: %s (available %d, requested %d)", e.Product, e.Available, e.Requested)
}

func (e OutOfStockError) Unwrap() error {
	return nil
}

type InsufficientFundsError struct {
	UserID  string
	Balance float64
	Cost    float64
}

func (e InsufficientFundsError) Error() string {
	return fmt.Sprintf("insufficient funds for user %s: balance %.2f, cost %.2f", e.UserID, e.Balance, e.Cost)
}

func (e InsufficientFundsError) Unwrap() error {
	return nil
}
func handleOrderError(err error) string {
	if errors.Is(err, ErrOrderNotFound) {
		return "Заказ не найден."
	}

	var se OutOfStockError
	if errors.As(err, &se) {
		return fmt.Sprintf("Товара %s нет в нужном количестве: запрошено %d, в наличии %d", se.Product, se.Requested, se.Available)
	}

	var ife InsufficientFundsError
	if errors.As(err, &ife) {
		return fmt.Sprintf("Недостаточно средств у пользователя %s: баланс %.2f, стоимость %.2f", ife.UserID, ife.Balance, ife.Cost)
	}
	return fmt.Sprintf("Неизвестная ошибка: %v", err)
}

func main() {
	errors := []error{
		ErrOrderNotFound,
		OutOfStockError{Product: "Телефон", Available: 5, Requested: 10},
		InsufficientFundsError{UserID: "user123", Balance: 1000.50, Cost: 1500.75},
		fmt.Errorf("обёрнутая ошибка: %w", ErrOrderNotFound),
		fmt.Errorf("обёрнутая ошибка: %w", OutOfStockError{Product: "Наушники", Available: 2, Requested: 3}),
	}

	for _, err := range errors {
		fmt.Println(handleOrderError(err))
	}
}
