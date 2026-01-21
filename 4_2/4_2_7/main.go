package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	var n int
	fmt.Scan(&n)

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {

	} // Пропускаем остаток строки после n

	words := make([]string, n)
	for i := 0; i < n; i++ {
		if scanner.Scan() {
			fmt.Println("Ошибка ввода: недостаточно строк")
			return
		}
		words[i] = scanner.Text()
	}

	result := winnerSlice(words)
	fmt.Println(result)
}

func winnerSlice(words []string) string {
	const Masha = "Маша"
	const Pasha = "Паша"

	for i := 1; i < len(words); i++ {
		prevWord := words[i-1]
		currWord := words[i]

		lastPrev := getLastRune(prevWord)
		firstCurr := getFirstRune(currWord)

		if unicode.ToLower(lastPrev) != unicode.ToLower(firstCurr) {
			if i%2 == 1 {
				return Masha // Паша ошибся → Маша победила
			} else {
				return Pasha // Маша ошиблась → Паша победил
			}
		}
	}

	// Все слова корректны → побеждает тот, кто сказал последнее слово
	if len(words)%2 == 1 {
		return Masha // Нечётное число слов → последнее слово сказала Маша
	} else {
		return Pasha // Чётное число слов → последнее слово сказал Паша
	}
}

func getFirstRune(s string) rune {
	for char := range s {
		return rune(char) // char уже rune
	}
	return 0
}

func getLastRune(s string) rune {
	r, _ := utf8.DecodeLastRuneInString(s)
	return r
}
