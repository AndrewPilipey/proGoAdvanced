package main

import (
	"errors"
	"fmt"
)

type Product struct {
	ID    int
	Name  string
	Price float64
}

func main() {
	var n int
	_, _ = fmt.Scan(&n)

	AllProduct := make([]Product, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&AllProduct[i].ID, &AllProduct[i].Name, &AllProduct[i].Price)
	}

	var targetID int
	_, _ = fmt.Scan(&targetID)

	product, err := FindProduct(AllProduct, targetID)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Printf("Товар: %s (%.2f руб.)\n", product.Name, product.Price)
	}
}

func FindProduct(products []Product, id int) (Product, error) {
	err := errors.New("product not found")

	for i := range products {
		if products[i].ID == id {
			return products[i], nil
		}
	}
	return Product{}, fmt.Errorf("товар с артикулом %d не найден: %w", id, err)
}
