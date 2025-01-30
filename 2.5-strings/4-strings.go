package string

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Check4String() {
	var text string
	text, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Println(removeEvenSymbols(strings.TrimSpace(text)))
}

func removeEvenSymbols(s string) string {
	rs := []rune(s)
	sb := strings.Builder{}
	for i := 1; i < len(rs); i += 2 {
		sb.WriteRune(rs[i])
	}

	return sb.String()
}
