package string

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func CheckString() {
	var text string
	text, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Println(isStringValid(strings.TrimSpace(text)))
}

func isStringValid(s string) string {
	if len(s) == 0 {
		return "Wrong"
	}

	rs := []rune(s)

	if unicode.IsUpper(rs[0]) && rs[len(rs)-1] == '.' {
		return "Right"
	}
	return "Wrong"
}
