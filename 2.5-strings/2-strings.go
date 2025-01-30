package string

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var text string
	text, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Println(isPalindrome(strings.TrimSpace(text)))
}

func isPalindrome(s string) string {
	rs := []rune(s)
	for i := 0; i < len(rs)/2; i++ {
		if rs[i] != rs[len(rs)-1-i] {
			return "Нет"
		}
	}

	return "Палиндром"
}
