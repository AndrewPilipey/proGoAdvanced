package main

import (
	"fmt"
)

type Student struct {
	FirstName string
	LastName  string
	Age       int
}

type School struct {
	Name     string
	Students []Student
}

func main() {
	fmt.Println("Здравствуйте! Введите название вашей школы:")
	var schoolName string
	_, _ = fmt.Scanln(&schoolName)

	school := School{Name: schoolName, Students: []Student{}}
	fmt.Printf("Школа \"%s\" успешно создана\n", school.Name)

	for {

		fmt.Println("\n --- Меню школ --- \n Выберите пункт 1-3:")
		fmt.Println("1. Посмотреть список учеников.")
		fmt.Println("2. Добавить нового ученика.")
		fmt.Println("3. Выйти из программы.")
		fmt.Println("9. Удалить ученика.")

		var choise int
		_, _ = fmt.Scanln(&choise)

		switch choise {
		case 1:
			school.PrintStudents()
		case 2:
			fmt.Println("Введите имя ученика: ")
			var firstName string
			_, _ = fmt.Scanln(&firstName)

			fmt.Println("Введите фамилию ученика: ")
			var lastName string
			_, _ = fmt.Scanln(&lastName)

			fmt.Println("Введите возраст ученика: ")
			var age int
			_, _ = fmt.Scanln(&age)

			student := Student{FirstName: firstName, LastName: lastName, Age: age}
			school = school.AddNewStudent(student)
		case 3:
			fmt.Println("До свидания!")
			return
		case 9:
			fmt.Println("Введите номер ученика")
			var number int
			_, _ = fmt.Scanln(&number)
			school.DeletingStudent(number)
		default:
			fmt.Println("Неверный пункт меню. Повторите ввод: ")
		}
	}
}

func (s School) PrintStudents() {
	if len(s.Students) == 0 {
		fmt.Printf("В школе \"%s\" пока нет учеников!\n", s.Name)
	} else {
		fmt.Printf("Список учеников:\n%-10s %-10s %-5s", "Имя", "Фамилия", "Возраст\n")
		for _, student := range s.Students {
			fmt.Printf("%-10s %-10s %-5d\n", student.FirstName, student.LastName, student.Age)
		}
	}
}

func (s School) AddNewStudent(Student Student) School {
	s.Students = append(s.Students, Student)
	fmt.Printf("Студент %s успешно добавлен в школу \"%s\".\n", Student.FirstName, s.Name)
	return s
}

func (s School) DeletingStudent(number int) School {
	if number < 1 || number > len(s.Students) {
		fmt.Println("Ошибка, неверный номер ученика. Допустимые значения: 1-", len(s.Students))
		return s
	}

	index := number - 1

	s.Students = append(s.Students[:index], s.Students[index+1:]...)
	fmt.Printf("Ученик N%d успешно удалён из школы \"%s\".\n", number, s.Name)
	return s
}
