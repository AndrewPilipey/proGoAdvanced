package main

import (
	"errors"
	"fmt"
)

func validateAge(age int) error {
	if age <= 0 {
		return errors.New("возраст должен быть положительным")
	} else if age < 18 {
		return errors.New("возраст должен быть не меньше 18")
	} else {
		return nil
	}
} //

func main() {
	ages := []int{-5, 16, 25}

	for _, age := range ages {
		err := validateAge(age)
		if err != nil {
			fmt.Printf("Возраст %d: ошибка - %s\n", age, err)
		} else {
			fmt.Printf("Возраст %d: подходит\n", age)
		}
	}
}
