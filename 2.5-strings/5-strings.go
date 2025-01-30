package string

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Check5String() {
	var text string
	text, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Println(removeDoubleSymbols(strings.TrimSpace(text)))
}

func removeDoubleSymbols(s string) string {
	sb := strings.Builder{}
	for _, ch := range s {
		if strings.Count(s, string(ch)) < 2 {
			sb.WriteRune(ch)
		}
	}

	return sb.String()
}
