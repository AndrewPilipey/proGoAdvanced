package main

import (
	"fmt"
)

type Cafe struct {
	Name     string
	Director Director
}

type Director struct {
	Name        string
	Managers    []Manager
	Departments []Department
}

type Manager struct {
	Name       string
	Department Department
}

type Department struct {
	Name  string
	Staff []Employee
}

type Employee struct {
	FullName string
	Position string
}

func main() {
	cafe := Cafe{}
	cafe.Name = "У Артура"

	director := Director{}
	director.Name = "Юрко В. В."
	director.Departments = []Department{}
	director.Managers = []Manager{}
	cafe.Director = director

	department := Department{}
	department.Name = "Коммерческий отдел"
	department.Staff = []Employee{}
	director.Departments = append(director.Departments, department)

	manager := Manager{}
	manager.Name = "Коршун А. С."
	manager.Department = department
	director.Managers = append(director.Managers, manager)

	admin := Employee{}
	admin.FullName = "Котько П. Д."
	admin.Position = "Администратор"
	department.Staff = append(department.Staff, admin)

	fmt.Printf("Кафе \"%s\"\n", cafe.Name)
	fmt.Printf("Директор: %s\n", director.Name)
	fmt.Printf("Менеджер: %s\n", manager.Name)
	fmt.Printf("Подразделение: \"%s\"\n", department.Name)
	fmt.Printf("%s: %s\n", admin.Position, admin.FullName)
}
