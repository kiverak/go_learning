package types

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func DoTypes3() {
	text, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil && err != io.EOF {
		fmt.Print("Error EOF")
	}

	text = strings.TrimSpace(text)
	arrText := strings.Split(text, ";")
	if len(arrText) != 2 {
		fmt.Println("Неверный формат ввода. Ожидается два числа, разделённые ';'")
		return
	}

	result := calculateDevisionFromStrings(arrText[0], arrText[1])
	fmt.Printf("%.4f\n", result)
}

func calculateDevisionFromStrings(str1, str2 string) float64 {
	num1 := extractNumFromUnformattedString(str1)
	num2 := extractNumFromUnformattedString(str2)
	return num1 / num2
}

func extractNumFromUnformattedString(str string) float64 {
	text := strings.ReplaceAll(str, " ", "")
	text = strings.ReplaceAll(text, ",", ".")
	num, _ := strconv.ParseFloat(text, 64)
	return num
}
