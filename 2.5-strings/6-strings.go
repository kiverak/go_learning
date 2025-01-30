package string

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

func Check6String() {
	var text string
	text, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	if validatePassword(strings.TrimSpace(text)) {
		fmt.Println("Ok")
	} else {
		fmt.Println("Wrong password")
	}
}

func validatePassword(s string) bool {
	if utf8.RuneCountInString(s) < 5 {
		return false
	}

	for _, ch := range s {
		if !(unicode.Is(unicode.Latin, ch) || unicode.Is(unicode.Arabic, ch) || unicode.IsNumber(ch)) {
			return false
		}
	}

	return true
}
