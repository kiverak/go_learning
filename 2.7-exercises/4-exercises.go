package exercises2

import (
	"fmt"
	"strconv"
	"strings"
)

func DoExercise4() {
	var text string
	_, err := fmt.Scan(&text)
	if err != nil {
		return
	}
	fmt.Println(powEveryFigure(text))
}

func powEveryFigure(text string) string {
	builder := strings.Builder{}

	for _, n := range text {
		num, _ := strconv.Atoi(string(n))
		builder.WriteString(strconv.Itoa(num * num))
	}

	return builder.String()
}
