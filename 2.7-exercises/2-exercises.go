package string

import (
	"fmt"
	"strings"
)

func DoExercise2() {
	var text string
	fmt.Scan(&text)
	fmt.Println(addAsteriscs(text))
}

func addAsteriscs(text string) string {
	runes := []rune(text)
	builder := strings.Builder{}

	for i, r := range runes {
		builder.WriteRune(r)
		if i != len(runes)-1 {
			builder.WriteRune('*')
		}
	}

	return builder.String()
}
