package main

import (
	"errors"
	"fmt"
	"strings"
)

type Applicantes struct {
	examScore     int
	documents     bool
	age           int
	requiredScore int //
}

func (a *Applicantes) CheckExam() error {
	if a.examScore < a.requiredScore {
		return fmt.Errorf("баллов недостаточно: %d < %d", a.examScore, a.requiredScore)
	}
	return nil
}

func (a Applicantes) CheckDocuments() error {
	if !a.documents {
		return errors.New("не все документы предоставлены")
	}
	return nil
}

func (a Applicantes) CheckAge() error {
	if a.age < 16 {
		return errors.New("возраст меньше минимального (16)")
	}
	if a.age > 35 {
		return errors.New("возраст больше максимального (35)")
	}
	return nil
}

func (a Applicantes) CheckAll() error {
	var errs []error

	if err := a.CheckExam(); err != nil {
		errs = append(errs, err)
	}

	if err := a.CheckDocuments(); err != nil {
		errs = append(errs, err)
	}

	if err := a.CheckAge(); err != nil {
		errs = append(errs, err)
	}

	return errors.Join(errs...)
}

func main() {
	var requiredScore int //прох.балл
	_, err := fmt.Scanln(&requiredScore)
	if err != nil {
		fmt.Println("Ошибка ввода проходного балла:", err)
		return
	}

	var n int //кол-во студентов
	_, err = fmt.Scanln(&n)
	if err != nil {
		fmt.Println("Ошибка ввода колтчесьва студентов:", err)
	}

	students := make([]Applicantes, n)

	for i := range students {
		var pass Applicantes
		pass.requiredScore = requiredScore
		fmt.Scanln(&pass.examScore, &pass.documents, &pass.age)
		students[i] = pass
	}

	for _, st := range students {
		err := st.CheckAll()
		if err != nil {
			fmt.Println("Проблемы:")
			problems := strings.Split(err.Error(), "\n")
			for _, problem := range problems {
				if problem != "" {
					fmt.Println("-", problem)
				}
			}
		} else {
			fmt.Println("Заявление OK")
		}
	}
}
