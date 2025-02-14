package types

import (
	"fmt"
	"strings"
	"unicode"
)

type Battery struct {
	Level uint8
}

func (b Battery) String() string {
	return fmt.Sprintf("[%10s]", strings.Repeat("X", int(b.Level)))
}

func DoInterfaces22() {
	// batteryForTest - не забывайте об имени

	var str string
	fmt.Scan(&str)
	batteryForTest := Battery{Level: getLevel(str)}
	fmt.Print(batteryForTest.String())
}

func getLevel(str string) uint8 {
	return uint8(strings.Count(str, `1`))
}

/********************************************************************************/

type customError uint

func (c customError) Error() string {
	return fmt.Sprintf("цифра, индекс %d", c)
}

func errorInString(str string) error {
	// Полезная работа со строкой проигнорирована
	for i, s := range str {
		if unicode.IsDigit(s) {
			return customError(i)
		}
	}
	return nil
}

func DoInterfaces2() {
	err := errorInString("string1string")
	if err != nil {
		fmt.Printf("Ошибка обработана: %v\n", err)
	}
	if cError, ok := err.(customError); ok {
		fmt.Printf("Контекст: %d\n", cError)
	}

	// Output:
	// Ошибка обработана: цифра, индекс 6
	// Контекст: 6
}
