package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
	"unicode/utf8"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() { // Проверяю успешность сканирования
		fmt.Println("Ошибка чтения n")
		return

	}
	n, _ := strconv.Atoi(scanner.Text())

	words := make([]string, n)
	for i := 0; i < n; i++ {
		if !scanner.Scan() { // Проверяю успешность сканирования
			fmt.Println("Ошибка ввода слов.")
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
				return Pasha // Маша ошибся → Паша победила
			} else {
				return Masha // Паша ошиблась → Маша победил
			}
		}
	}

	// Все слова корректны → побеждает тот, кто сказал последнее слово
	if len(words)%2 == 0 {
		return Masha // Чётное число слов → последнее слово сказал Паша
	} else {
		return Pasha // Нечётное число слов → последнее слово сказала Маша
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
