package main

import "fmt"

const (
	Low = iota + 1
	Middle
	High
	Critical
)

func main() {
	var (
		p  int
		p1 int
		p2 int
	)

	_, _ = fmt.Scan(&p, &p1, &p2)
	fmt.Println(definePriority(p))
	fmt.Println(isHigher(p1, p2))
}

func isHigher(p1, p2 int) bool {
	if p1 > p2 {
		return true
	}
	return false
}

func definePriority(p int) string {
	switch p {
	case 0:
		return "Low"
	case 1:
		return "Medium"
	case 2:
		return "High"
	case 3:
		return "Critical"
	default:
		return "Неизвестный приоритет"
	}
}
