package main

import (
	"fmt"
)

type Fraction struct {
	Numerator   int
	Denominator int
}

func main() {
	fmt.Println("Здесь Вы сможете производить операции с простыми дробями")

	fmt.Println("Задайте значения простой дроби #1:")
	var numerator, denominator int
	fmt.Print("Числитель: ")
	_, err := fmt.Scan(&numerator)
	if err != nil {
		fmt.Println("Ошибка ввода числителя!")
		return
	}
	fmt.Print("Знаменатель: ")
	_, err = fmt.Scan(&denominator)
	if err != nil {
		fmt.Println("Ошибка ввода знаменателя!")
		return
	}
	if denominator == 0 {
		fmt.Println("Ошибка: знаменатель не может быть нулём!")
		return
	}
	fraction1 := Fraction{Numerator: numerator, Denominator: denominator}
	fraction1.Print() //

	fmt.Println("Выберите второй операнд:")
	fmt.Printf("%-5s %-5s %-5s", "1.Целое число", "2.Простая дробь", "0.Выход\n")

	var answer1 string
	_, err = fmt.Scan(&answer1)
	if err != nil {
		fmt.Println("Ошибка выбора операнда.")
		return
	}
	switch answer1 {
	case "1":
		fmt.Println("Введите число: ")
		var number2 int
		_, err = fmt.Scan(&number2)
		if err != nil {
			fmt.Println("Ошибка ввода числа!")
			return
		}
	point1:
		fmt.Println("Выберите операцию: ")
		fmt.Printf("%-5s %-5s %-5s %-5s %-5s", "1.Сложение", "2.Вычитание", "3.Умножение", "4.Деление", "0.Выход\n")
		var answer2 string
		_, err = fmt.Scan(&answer2)
		if err != nil {
			fmt.Println("Ошибка выбора операции.")
			goto point1
		}

		switch answer2 {
		case "1":
			result := fraction1.SumInt(number2)
			result.Print()
		case "2":
			result := fraction1.DifferenceInt(number2)
			result.Print()
		case "3":
			result := fraction1.MultiplyInt(number2)
			result.Print()
		case "4":
			if number2 == 0 {
				fmt.Println("На ноль делить нельзя.")
				goto point1
			}
			result := fraction1.DivideInt(number2)
			result.Print()
		case "0":
			return
		default:
			fmt.Println("Что-то пошло не так. Повторите операцию.")
			goto point1

		}

	case "2":
		fmt.Println("Задайте значения простой дроби #2:")
		var numerator2, denominator2 int
		fmt.Print("Числитель: ")
		_, err = fmt.Scan(&numerator2)
		if err != nil {
			fmt.Println("Ошибка ввода числителя!")
			return
		}
		fmt.Print("Знаменатель: ")
		_, err = fmt.Scan(&denominator2)
		if err != nil {
			fmt.Println("Ошибка ввода знаменателя!")
			return
		}
		if denominator2 == 0 {
			fmt.Println("Ошибка: знаменатель не может быть нулём!")
			return
		}
		fraction2 := Fraction{Numerator: numerator2, Denominator: denominator2}
		fraction2.Print() //
	point2:
		fmt.Println("Выберите операцию:")
		fmt.Printf("%-5s %-5s %-5s %-5s %-5s", "1.Сложение", "2.Вычитание", "3.Умножение", "4.Деление", "0.Выход\n")
		var answer3 string
		_, err = fmt.Scan(&answer3)
		if err != nil {
			fmt.Println("Ошибка выбора операции.")
			goto point2
		}

		switch answer3 {
		case "1":
			result := fraction1.Sum(fraction2)
			result.Print()
		case "2":
			result := fraction1.Difference(fraction2)
			result.Print()

		case "3":
			result := fraction1.Multiply(fraction2)
			result.Print()
		case "4":
			result := fraction1.Divide(fraction2)
			result.Print()
		case "0":
			return
		default:
			fmt.Println("Что-то пошло не так. Повторите операцию.")
			goto point2
		}
	case "0":
		return
	}
}

// метод вывода
func (f Fraction) Print() {
	fmt.Printf("Ответ: %d/%d\n", f.Numerator, f.Denominator)
}

// метод округления дроби перед выводом результата
func (f Fraction) reduce() Fraction {
	gcd := f.gcd() //Greatest Common Divisor
	return Fraction{
		Numerator:   f.Numerator / gcd,
		Denominator: f.Denominator / gcd,
	}
}

// метод gcd (Greatest common divisor)
// - нахождение наибольшего общего делителя по алгоритму Евклида
func (f Fraction) gcd() int {
	a, b := f.Numerator, f.Denominator
	if b == 0 { // Если знаменатель 0, НОД = числитель по определению
		return abs(a)
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// методы операций. Действия с дробями
func (f Fraction) Sum(other Fraction) Fraction {
	commonDenominator := f.Denominator * other.Denominator
	resultNumerator := f.Numerator*other.Denominator + other.Numerator*f.Denominator
	result := Fraction{Numerator: resultNumerator, Denominator: commonDenominator}
	result = result.reduce()
	return result
}

func (f Fraction) Difference(other Fraction) Fraction {
	commonDenominator := f.Denominator * other.Denominator
	resultNumerator := f.Numerator*other.Denominator - other.Numerator*f.Denominator
	result := Fraction{Numerator: resultNumerator, Denominator: commonDenominator}
	result = result.reduce()
	return result
}

func (f Fraction) Multiply(other Fraction) Fraction {
	result := Fraction{Numerator: f.Numerator * other.Numerator, Denominator: f.Denominator * other.Denominator}
	result = result.reduce()
	return result
}

func (f Fraction) Divide(other Fraction) Fraction {
	result := Fraction{Numerator: f.Numerator * other.Denominator, Denominator: f.Denominator * other.Numerator}
	result = result.reduce()
	return result
}

// методы операций. Действия с целыми числами
func (f Fraction) SumInt(number int) Fraction {
	otherFraction := Fraction{Numerator: number, Denominator: 1}
	result := f.Sum(otherFraction)
	result = result.reduce()
	return result
}

func (f Fraction) DifferenceInt(number int) Fraction {
	otherFraction := Fraction{Numerator: number, Denominator: 1}
	result := f.Difference(otherFraction)
	result = result.reduce()
	return result
}

func (f Fraction) MultiplyInt(number int) Fraction {
	otherFraction := Fraction{Numerator: number, Denominator: 1}
	result := f.Multiply(otherFraction)
	result = result.reduce()
	return result
}

func (f Fraction) DivideInt(number int) Fraction {
	otherFraction := Fraction{Numerator: number, Denominator: 1}
	result := f.Divide(otherFraction)
	result = result.reduce()
	return result
}

// в целях доп защиты от деления на ноль
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
