package main

import (
	"fmt"
	"strings"
)

const (
	Read    = 1 << iota // битовый сдвиг 1 << 0 = 1 (001)
	Write               // 1 << 1 = 2 (010)
	Execute             //1 << 2 = 4 (100)
)

func main() {
	var accessCode int
	_, _ = fmt.Scan(&accessCode)

	fmt.Println(readAccessCode(accessCode))

}

func readAccessCode(accessCode int) string {
	var result []string

	if accessCode&Read != 0 {
		result = append(result, "можно читать")
	}
	if accessCode&Write != 0 {
		result = append(result, "можно писать")
	}
	if accessCode&Execute != 0 {
		result = append(result, "можно выполнять")
	}

	if len(result) == 0 {
		return ""
	}

	return strings.Join(result, "\n")
}
