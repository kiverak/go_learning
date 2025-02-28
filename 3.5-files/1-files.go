package files

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func DoFiles1() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var sum = 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		n, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		sum += n
	}
	if err := scanner.Err(); err == nil || err == io.EOF {
		writer.WriteString(strconv.Itoa(sum))
	} else {
		panic(err)
	}
}
