package main

import (
	"errors"
	"fmt"
)

func main() {
	var n int //кол-во транзакций
point0:
	_, err := fmt.Scanln(&n)
	if err != nil {
		fmt.Println("Ошибка ввода количества транзакций. Повторите попытку.")
		goto point0
	}

	for i := 0; i < n; i++ {
		var id, amount int
		_, _ = fmt.Scanln(&id, &amount)

		err := processTransaction(id, amount)
		if err != nil {
			fmt.Printf("Ошибка: %v\n", err)
		} else {
			fmt.Printf("Транзакция #%d успешно завершена\n", id)
		}
	}

}

func processTransaction(id int, amount int) error {
	hasError := false

	defer func() {
		if hasError {
			fmt.Printf("Транзакция #%d: ОТКАТ\n", id)
			hasError = false
		}
	}()

	if amount <= 0 {
		hasError = true
		return errors.New("сумма должна быть положительной")
	} else if amount > 10_000 {
		hasError = true
		return errors.New("сумма превышает лимит")
	} else {
		fmt.Printf("Транзакция #%d: валидация пройдена\n", id)
	}

	fmt.Printf("Транзакция #%d: средства списаны\n", id)

	if amount%1000 == 0 {
		hasError = true
		return errors.New("ошибка подтверждения транзакции")
	} else {
		fmt.Printf("Транзакция #%d: подтверждена\n", id)

	}

	return nil
}
