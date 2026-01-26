package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type roaming struct {
	name string
	city string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nStr := scanner.Text()
	n, err := strconv.Atoi(nStr)
	if err != nil {
		fmt.Println("Ошибка")
		return
	}

	var people []roaming

	for i := 0; i < n; i++ {
		scanner.Scan()
		text := scanner.Text()
		words := strings.Fields(text)

		if len(words) < 2 {
			fmt.Printf("Ошибка: в строке недостаточно слов")
			continue
		}

		person := roaming{
			name: words[0],
			city: words[1],
		}
		people = append(people, person)
	}

	for _, p := range people {
		message := fmt.Sprintf("Добрый день, %s! Рады приветствовать вас в городе %s!", p.name, p.city)
		fmt.Println(message)
	}

}
