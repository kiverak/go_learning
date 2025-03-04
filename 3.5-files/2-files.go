package files

import (
	"archive/zip"
	"encoding/csv"
	"fmt"
	"strings"
)

func DoFiles2() {
	r, err := zip.OpenReader("./files/3.5.2/task.zip")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer r.Close()

	for _, file := range r.File {
		// Open the file within the zip archive
		if strings.Contains(file.Name, ".txt") {
			rc, err := file.Open()
			if err != nil {
				fmt.Println(err)
				continue
			}
			defer rc.Close()

			reader := csv.NewReader(rc)
			records, err := reader.ReadAll()
			if err != nil {
				fmt.Printf("reading file error = %s", err)
				continue
			}

			if len(records) == 10 {
				fmt.Println(records[4][2])
			}
		}
	}
}
