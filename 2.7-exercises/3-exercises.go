package exercises2

import (
	"fmt"
	"unicode"
)

func DoExercise3() {
	var text string
	_, err := fmt.Scan(&text)
	if err != nil {
		return
	}
	fmt.Println(findMaxFigure(text))
}

func findMaxFigure(text string) string {
	runes := []rune(text)
	var max rune
	for _, n := range runes {
		if unicode.IsDigit(n) {
			max = n
			break
		}
	}
	for _, n := range runes {
		if unicode.IsDigit(n) && n > max {
			max = n
		}
	}

	return string(max)
}
