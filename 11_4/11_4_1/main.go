package main

import (
	"errors"
	"fmt"
)

func safeDivide(a, b int) (result int, err error) {

	defer func() {
		if r := recover(); r != nil {
			//fmt.Println("Паника перехвачена")
			result = 0
			err = errors.New("деление на ноль")

		}
	}()
	return a / b, nil
}

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scan(&a, &b)

		res, err := safeDivide(a, b)
		if err != nil {
			fmt.Println("Ошибка:", err)
		} else {
			fmt.Println("Результат:", res)
		}
	}
}
