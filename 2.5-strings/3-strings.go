package string

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Check3String() {
	var text, es string
	text, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	es, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Println(getFirstContaining(strings.TrimSpace(text), strings.TrimSpace(es)))
}

func getFirstContaining(s, es string) int {
	rs := []rune(s)
	rsEs := []rune(es)

	for i := 0; i < len(rs)-len(rsEs); i++ {
		for j := 0; j < len(rsEs); j++ {
			if rs[i+j] != rsEs[j] {
				break
			} else if j == len(rsEs)-1 {
				return i - j
			}
		}
	}

	return -1
}
