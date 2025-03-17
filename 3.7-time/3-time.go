package time

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func DoTime3() {
	reader := bufio.NewReader(io.Reader(os.Stdin))
	time1, err := reader.ReadString(',')
	if err != nil {
		fmt.Println("Error reading input 1:", err)
		return
	}
	time1 = strings.Trim(time1, ",")

	time2, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		fmt.Println("Error reading input 2:", err)
		return
	}
	time2 = strings.TrimSpace(time2)

	timeParsed1, err := time.Parse("02.01.2006 15:04:05", time1)
	if err != nil {
		fmt.Println("Error parsing time1:", err)
		return
	}
	timeParsed2, err := time.Parse("02.01.2006 15:04:05", time2)
	if err != nil {
		fmt.Println("Error parsing time2:", err)
		return
	}

	sub := timeParsed1.Sub(timeParsed2)

	fmt.Print(sub.Abs())
}
