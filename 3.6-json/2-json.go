package files

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type global struct {
	GlobalId int `json:"global_id"`
}

type muStruct2 struct {
	Global []global `json:"global"`
}

func DoJson2() {
	urlDownload := "https://github.com/semyon-dev/stepik-go/blob/master/work_with_json/data-20190514T0100.json"
	resp, err := http.Get(urlDownload)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading from url:", err)
		return
	}

	var gl global
	buf := new(bytes.Buffer)
	dec := json.NewDecoder(buf)
	sum := 0
	for _, bytes := range data {
		buf.WriteByte(bytes)
		dec.Decode(&gl)
		sum += gl.GlobalId
		buf.Reset()
	}

	fmt.Print(strconv.Itoa(sum))
}
