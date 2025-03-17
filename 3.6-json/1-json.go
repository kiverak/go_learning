package json

import (
	"encoding/json"
	"fmt"
	"os"
)

type myStruct struct {
	ID       int    `json:"ID"`
	Number   string `json:"Number"`
	Year     int    `json:"Year"`
	Students []struct {
		LastName   string `json:"LastName,omitempty"`
		FirstName  string `json:"FirstName,omitempty"`
		MiddleName string `json:"MiddleName,omitempty"`
		Birthday   string `json:"Birthday,omitempty"`
		Address    string `json:"Address,omitempty"`
		Phone      string `json:"Phone,omitempty"`
		Rating     []int  `json:"Rating,omitempty"`
	} `json:"Students"`
}

type average struct {
	Average float32
}

func DoJson1() {
	//data, err := io.ReadAll(os.Stdin)
	fileName := "./files/3.6.1/test_file.txt"
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading from stdin:", err)
		return
	}

	var s myStruct

	if err := json.Unmarshal(data, &s); err != nil {
		fmt.Println(err)
		return
	}

	count := 0
	for _, student := range s.Students {
		count += len(student.Rating)
	}

	avg := float32(count) / float32(len(s.Students))

	// Output:
	output := average{Average: avg}
	data, err = json.MarshalIndent(output, "", "\t")
	fmt.Print(string(data))
}
