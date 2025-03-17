package time

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func DoTime2() {
	reader := bufio.NewReader(io.Reader(os.Stdin))
	timeString, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	timeString = strings.TrimSpace(timeString)

	timeParsed, err := time.Parse(time.DateTime, timeString)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return
	}

	dateTime := time.Date(timeParsed.Year(), timeParsed.Month(), timeParsed.Day(), 13, 0, 0, 0, timeParsed.Location())

	if timeParsed.After(dateTime) {
		timeParsed = timeParsed.AddDate(0, 0, 1)
	}

	fmt.Print(timeParsed.Format(time.DateTime))
}
