package string

import (
	"errors"
	"fmt"
)

func LearnErrors() {
	var a, b int
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println("ошибка")
		return
	}

	i, err := divide(a, b)
	if err != nil {
		fmt.Println("ошибка")
	} else {
		fmt.Println(i)
	}
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("ошибка")
	}

	return a / b, nil
}
