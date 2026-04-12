package main

import "fmt"

const (
	_ = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func main() {
	var day int
	_, _ = fmt.Scan(&day)

	fmt.Println(dayName(day))

}

func dayName(day int) string {
	switch day {
	case Monday:
		return "Понедельник"
	case Tuesday:
		return "Вторник"
	case Wednesday:
		return "Среда"
	case Thursday:
		return "Четверг"
	case Friday:
		return "Пятница"
	case Saturday:
		return "Суббота"
	case Sunday:
		return "Воскресенье"
	default:
		return "Некорректный день"
	}
}
