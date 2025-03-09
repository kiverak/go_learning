package files

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func DoFiles3() {
	file, err := os.Open("./files/3.5.3/task.data")
	if err != nil {
		fmt.Println("Unable to open file:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	counter := 0
	for {
		line, err := reader.ReadString(';')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				return
			}
		}
		counter++
		line = strings.TrimSuffix(line, ";")
		if line == "0" {
			fmt.Println(counter)
			break
		}
	}
}
