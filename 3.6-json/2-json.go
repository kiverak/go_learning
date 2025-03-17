package json

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type global struct {
	GlobalId int `json:"global_id"`
}

func DoJson2() {
	//urlDownload := "https://github.com/semyon-dev/stepik-go/blob/master/work_with_json/data-20190514T0100.json"
	//resp, err := http.Get(urlDownload)
	//if err != nil {
	//	return
	//}
	//defer resp.Body.Close()
	//
	//file, err := io.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Println("Error reading from url:", err)
	//	return
	//}

	fileName := "./files/3.6.2/data-20190514T0100.json"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error opening file:", err)
		return
	}
	defer file.Close()

	// Decode the JSON content into an array of structs
	var dataArray []global
	err = json.NewDecoder(file).Decode(&dataArray)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	sum := 0
	for _, gl := range dataArray {
		sum += gl.GlobalId
	}

	fmt.Println(sum)
}

func DoJson3() {
	fileName := "./files/3.6.2/data-20190514T0100.json"
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading from file:", err)
		return
	}

	// Remove square brackets and split into individual JSON objects
	data = bytes.TrimPrefix(data, []byte("["))
	data = bytes.TrimSuffix(data, []byte("]"))
	lines := bytes.Split(data, []byte("},{"))

	sum := 0
	for _, line := range lines {
		// Add back the curly braces for each JSON object
		if !bytes.HasPrefix(line, []byte("{")) {
			line = append([]byte("{"), line...)
		}
		if !bytes.HasSuffix(line, []byte("}")) {
			line = append(line, []byte("}")...)
		}

		var gl global
		err := json.Unmarshal([]byte(line), &gl)
		if err != nil {
			fmt.Println("Error decoding JSON:", err, string(line))
		}
		sum += gl.GlobalId
	}
	fmt.Println(strconv.Itoa(sum))
}

func DoJson4() {
	fileName := "./files/3.6.2/data-20190514T0100.json"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a JSON decoder
	dec := json.NewDecoder(file)

	// Read the opening bracket
	token, err := dec.Token()
	if err != nil {
		fmt.Println("Error reading JSON:", err)
		return
	}
	if token != json.Delim('[') {
		fmt.Println("Expected opening bracket")
		return
	}

	sum := 0

	// Decode each element in the array
	for dec.More() {
		var gl global
		err := dec.Decode(&gl)
		if err != nil {
			fmt.Println("Error decoding JSON element:", err)
			return
		}
		sum += gl.GlobalId
	}

	// Read the closing bracket
	token, err = dec.Token()
	if err != nil {
		fmt.Println("Error reading JSON:", err)
		return
	}
	if token != json.Delim(']') {
		fmt.Println("Expected closing bracket")
		return
	}

	fmt.Println(sum)
}
