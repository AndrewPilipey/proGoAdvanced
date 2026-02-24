package main

import (
	"fmt"
)

func main() {
	sliceNum := make([]int, 4)

	for i := range sliceNum {
		_, _ = fmt.Scanln(&sliceNum[i])
	}
	//fmt.Println(sliceNum)

	for i := range sliceNum {
		x, err := grade(sliceNum[i])
		if err != nil {
			fmt.Printf("%d: ошибка — %s", sliceNum[i], err)
		} else {
			fmt.Printf("%d: %s\n", sliceNum[i], x)
		}
	}

}

func grade(score int) (string, error) {
	switch {
	case score >= 90 && score <= 100:
		return "Отлично", nil
	case score >= 70 && score <= 89:
		return "Хорошо", nil
	case score >= 50 && score <= 69:
		return "Удовлетворительно", nil
	case score >= 0 && score <= 49:
		return "Неудовлетворительно", nil
	default:
		return "", fmt.Errorf("баллы %d вне допустимого диапазона (0-100)\n", score)

	}
}
