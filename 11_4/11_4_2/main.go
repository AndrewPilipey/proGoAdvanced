package main

import (
	"errors"
	"fmt"
)

func safeGetElement(slice []int, index int) (value int, err error) {
	defer func() {
		if r := recover(); r != nil {
			value = 0
			err = errors.New("индекс за пределами слайса")
		}
	}()
	return slice[index], nil
}

func main() {
	numbers := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var index int
		fmt.Scan(&index)

		val, err := safeGetElement(numbers, index)
		if err != nil {
			fmt.Println("Ошибка:", err)
		} else {
			fmt.Println("Значение:", val)
		}
	}
}
