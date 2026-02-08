package main

import (
	"fmt"
)

type Student struct {
	FirstName string
	LastName  string
}

type School struct {
	Name     string
	Students []Student
}

func (s School) PrintStudents() {
	if len(s.Students) == 0 {
		fmt.Printf("В школе \"%s\" пока нет учеников!\n", s.Name)
	} else {
		fmt.Printf("%-10s %-10s\n", "Имя", "Фамилия")
		for _, student := range s.Students {
			fmt.Printf("%-10s %-10s\n", student.FirstName, student.LastName)
		}
	}
}

func (s School) AddNewStudent(student Student) School {
	s.Students = append(s.Students, student)
	fmt.Printf("Студент %s успешно добавлен в школу \"%s\".\n", student.FirstName, s.Name)
	return s
}

func main() {
	fmt.Println("Здравствуйте! Введите название вашей школы")
	var schoolName string
	fmt.Scanln(&schoolName)

	school := School{Name: schoolName, Students: []Student{}}
	fmt.Printf("Школа %s успешно создана\n", school.Name)

	for {
		fmt.Printf("Хотите посмотреть список учеников школы %s? Введите да или нет.\n", school.Name)
		var userAnswer string
		fmt.Scanln(&userAnswer)
		if userAnswer == "да" {
			school.PrintStudents()
		}

		fmt.Printf("Хотите добавить нового ученика в школу %s? Введите да или нет.\n", school.Name)
		fmt.Scanln(&userAnswer)
		if userAnswer == "да" {
			fmt.Println("Введите имя ученика")
			var firstName string
			fmt.Scanln(&firstName)

			fmt.Println("Введите фамилию ученика")
			var lastName string
			fmt.Scanln(&lastName)

			student := Student{FirstName: firstName, LastName: lastName}
			school = school.AddNewStudent(student)
		}
	}
}
