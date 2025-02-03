package types

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func DoTypes2() {
	var str1, str2 string
	fmt.Scan(&str1, &str2)

	fmt.Println(adding(str1, str2))
}

func adding(str1, str2 string) int64 {

	str1 = clean(str1)
	str2 = clean(str2)

	num1, err := strconv.ParseInt(str1, 10, 64)
	if err != nil {
		panic(err)
	}
	num2, err := strconv.ParseInt(str2, 10, 64)
	if err != nil {
		panic(err)
	}

	return num1 + num2
}

func clean(str string) string {
	runes := []rune(str)

	builder := strings.Builder{}
	for _, r := range runes {
		if unicode.IsDigit(r) {
			builder.WriteRune(r)
		}
	}

	return builder.String()
}
