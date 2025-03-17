package time

import (
	"fmt"
	"time"
)

func DoTime1() {
	var timeString string
	_, err := fmt.Scan(&timeString)
	if err != nil {
		return
	}

	timeParsed, err := time.Parse(time.RFC3339, timeString)
	if err != nil {
		return
	}

	fmt.Print(timeParsed.Format(time.UnixDate))
}
