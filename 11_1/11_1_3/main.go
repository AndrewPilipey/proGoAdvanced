package main

import (
	"errors"
	"fmt"
)

func main() {
	var N int
	_, err := fmt.Scanln(&N)
	if err != nil {
		fmt.Println("ошибка ввода количества имён")
		return
	}
	slice := make([]string, N)
	for i := 0; i < N; i++ {
		//fmt.Println("Введите Ваше имя:")
		_, err = fmt.Scanln(&slice[i])
		if err != nil {
			slice[i] = ""
		}

	}

	for _, name := range slice {
		greet, err := greetUser(name)
		if err != nil {
			fmt.Printf("Ошибка: %v\n", err)
		} else {
			fmt.Printf("%s\n", greet)
		}
	}
}

func greetUser(name string) (string, error) {

	if len([]rune(name)) == 0 {
		return "", errors.New("имя не может быть пустым")
	} else if len([]rune(name)) < 2 {
		return "", fmt.Errorf("имя %s слишком короткое", name)
	} else {
		return fmt.Sprintf("Привет, %s!", name), nil
	}
}
