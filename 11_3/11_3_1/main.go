package main

import "fmt"

func main() {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		return
	}

	for i := 0; i < n; i++ {
		var id, value int
		fmt.Scanln(&id, &value)

		err := processData(id, value)
		if err != nil {
			fmt.Printf("Ошибка: %s\n", err)
		} else {
			fmt.Printf("Данные #%d обработаны успешно\n", id)
		}
	}

}

func processData(id int, value int) error {
	defer fmt.Printf("Обработка данных #%d завершена\n", id)
	switch {
	case value < 0:
		return fmt.Errorf("%v", "отрицательное значение")
	case value == 0:
		return fmt.Errorf("%v", "нулевое значение")
	case value > 100:
		return fmt.Errorf("%v", "значение превышает лимит")

	}
	return nil
}
