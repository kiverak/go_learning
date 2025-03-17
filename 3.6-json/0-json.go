package json

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func DoJson0() {
	type myStruct struct {
		Name   string
		Age    int
		Status bool
		Values []int
	}

	s := myStruct{
		Name:   "John Connor",
		Age:    35,
		Status: true,
		Values: []int{15, 11, 37},
	}
	fmt.Printf("%v", s)
	fmt.Println()

	data, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s", data)
	fmt.Println()

	data, err = json.MarshalIndent(s, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s", data)
	fmt.Println()

	var s1 myStruct

	if err := json.Unmarshal(data, &s1); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%v", s)
	fmt.Println()
	fmt.Println()

	// validate JSON

	type user struct {
		Name     string
		Email    string
		Status   bool
		Language []byte
	}

	m := user{Name: "John Connor", Email: "test email"}

	data, _ = json.Marshal(m)

	data = bytes.Trim(data, "{") // испортим json удалив '{'

	// функция json.Valid возвращает bool, true - если json правильный
	if !json.Valid(data) {
		fmt.Println("invalid json!") // вывод: invalid json!
	}

	fmt.Printf("%s", data)
}
