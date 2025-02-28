package files

import (
	"fmt"
	"os"
	"strconv"
)

func DoFiles01() {
	for i := 1; i <= 3; i++ {
		file, err := os.Create(strconv.Itoa(i) + "text.txt")
		if err != nil {
			panic(err)
		}
		fmt.Println(file.Stat())
		//defer file.Close()
		file.Close()
	}
	os.Rename("2text.txt", "4text.txt")
	for i := 4; i >= 2; i-- {
		os.Remove(strconv.Itoa(i) + "text.txt")
	}
}
