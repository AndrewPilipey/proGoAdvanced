package main

import (
	"fmt"
)

type ageError struct {
	age     int
	message string
}

func (e ageError) Error() string {
	return fmt.Sprintf("возраст %d: %s", e.age, e.message)
}

func ValidateAge(age int) error {
	switch {

	case age < 0:
		return ageError{age: age, message: "возраст не может быть отрицательным"}
	case age == 0:
		return ageError{age: age, message: "возраст не может быть равен 0"}
	case age > 150:
		return ageError{age: age, message: "возраст не может быть больше 150"}

	}
	return nil
}

func main() {
	ages := []int{-5, 0, 200, 25}
	for _, age := range ages {
		err := ValidateAge(age)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Возраст %d корректен\n", age)
		}
	}
}
