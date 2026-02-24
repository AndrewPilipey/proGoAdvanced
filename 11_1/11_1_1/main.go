package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("деление на ноль")
	}
	res := a / b
	return res, nil
}

func main() {
	res, err := divide(10, 2)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("10 / 2 =", res)
	}

	res, err = divide(10, 0)
	if err != nil {
		fmt.Println("Ошибка:", err)
	}
}
