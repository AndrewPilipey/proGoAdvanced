package main

import (
	"fmt"
)

type Student interface {
	Name() string
	ListeningLection()
	DoHomework(exercise string)
	PrepareForExam()
}

type StudentStruct struct {
	name string
}

func (s StudentStruct) Name() string {
	return s.name
}

func (s StudentStruct) ListeningLection() {
	fmt.Println("ListeningLection")
}

func (s StudentStruct) DoHomework(exercise string) {
	fmt.Println("DoHomework")
}

func (s StudentStruct) PrepareForExam() {
	fmt.Println("PrepareForExam")
}

func main() {
	var studentValue Student = StudentStruct{name: "Josef"}

	// Присваиваем интерфейсу Student
	//var student Student = studentValue

	fmt.Println(studentValue.Name())
	studentValue.ListeningLection()
	studentValue.DoHomework("123")
	studentValue.PrepareForExam()
}
