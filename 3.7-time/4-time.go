package time

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

const now = 1589570165

func DoTime4() {
	reader := bufio.NewReader(io.Reader(os.Stdin))
	minutesStr, err := reader.ReadString(' ')
	if err != nil {
		fmt.Println("Error reading minutes:", err)
		return
	}
	minutesStr = strings.TrimSpace(minutesStr)
	minutes, err := strconv.Atoi(minutesStr)
	if err != nil {
		fmt.Println("Error parsing minutes:", err)
		return
	}

	_, err = reader.ReadString(' ')
	if err != nil {
		fmt.Println("Error reading text:", err)
		return
	}

	secondsStr, err := reader.ReadString(' ')
	if err != nil {
		fmt.Println("Error reading seconds:", err)
		return
	}
	secondsStr = strings.TrimSpace(secondsStr)
	seconds, err := strconv.Atoi(secondsStr)
	if err != nil {
		fmt.Println("Error parsing seconds:", err)
		return
	}

	duration := time.Duration(minutes*60000000000 + seconds*1000000000)
	unixTime := time.Unix(now, duration.Nanoseconds()).In(time.UTC)

	fmt.Print(unixTime.Format(time.UnixDate))
}
