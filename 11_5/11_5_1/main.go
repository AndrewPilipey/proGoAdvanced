package main

import (
	"errors"
	"fmt"
)

type Student struct {
	ID   int
	Name string
	Age  int
}

func main() {
	var n int
	_, _ = fmt.Scan(&n)

	students := make([]Student, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&students[i].ID, &students[i].Name, &students[i].Age)
	}

	var targetID int
	fmt.Scan(&targetID)

	student, err := FindStudent(students, targetID)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Printf("Найден студент: %s (%d лет)\n", student.Name, student.Age)
	}
}

func FindStudent(Students []Student, id int) (Student, error) {
	for i := range Students {
		if Students[i].ID == id {
			return Students[i], nil
		}
	}
	err := errors.New("student not found")
	return Student{}, fmt.Errorf("студент с ID %d не найден: %w", id, err)
}
