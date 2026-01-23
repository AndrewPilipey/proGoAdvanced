package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fileName := scanner.Text()
	scanner.Scan()
	ext := scanner.Text()

	if !isValidExtention(ext) {
		fmt.Println("Корректное расширение файла не найдено")
		return
	}

	lastDot := strings.LastIndex(fileName, ".")
	if lastDot == -1 {
		fmt.Println("Корректное расширение файла не найдено")
		return
	}

	oldExt := fileName[lastDot+1:]
	if oldExt == "" {
		fmt.Println("Корректное расширение файла не найдено")
		return
	}

	for _, r := range oldExt {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			fmt.Println("Корректное расширение файла не найдено")
			return
		}
	}

	newFileName := fileName[:lastDot] + ext
	fmt.Println(newFileName)
}

func isValidExtention(ext string) bool {
	if !strings.HasPrefix(ext, ".") || len(ext) == 1 {
		return false
	}
	for _, j := range ext[1:] {
		if !unicode.IsLetter(j) && !unicode.IsDigit(j) {
			return false
		}
	}
	return true
}
