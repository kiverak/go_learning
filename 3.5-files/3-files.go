package files

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
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

func DoFiles31() {
	urlDownload := "https://raw.githubusercontent.com/semyon-dev/stepik-go/master/work_with_files/task_sep_1/task.data"
	resp, err := http.Get(urlDownload)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	r := csv.NewReader(resp.Body)
	r.Comma = ';'
	lines, _ := r.ReadAll()
	for index, val := range lines[0] {
		if val == "0" {
			fmt.Println("Result:", index+1)
		}
	}
}
