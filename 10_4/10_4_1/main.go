package main

import (
	"fmt"
)

func main() {
	fmt.Println(ProcessData("привет"))
	fmt.Println(ProcessData(100))
	fmt.Println(ProcessData(3.14159))
	fmt.Println(ProcessData([]string{"a", "b"}))
	fmt.Println(ProcessData(nil))
	fmt.Println(ProcessData(true))
}

func ProcessData(data any) string {
	switch t := data.(type) {
	case string:
		return fmt.Sprintf("Строка: %s", t)
	case int:
		return fmt.Sprintf("Целое число: %d", t)
	case float64:
		return fmt.Sprintf("Дробное число: %.2f", t)
	case []string:
		return fmt.Sprintf("Срез строк: %d", len(t))
	default:
		return "Неизвестный тип"

	}

}
