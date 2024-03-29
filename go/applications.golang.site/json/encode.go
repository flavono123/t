package main

import (
	"encoding/json"
	"fmt"
)

type Member struct {
	Name   string
	Age    int
	Active bool
}

func main() {
	mem := Member{"Alex", 10, true}

	jsonBytes, err := json.Marshal(mem)
	if err != nil {
		panic(err)
	}

	jsonString := string(jsonBytes)

	fmt.Println(jsonString)

	m := make(map[string]interface{})
	m["a"] = 1

	jsonBytes, err = json.Marshal(m)
	if err != nil {
		panic(err)
	}

	jsonString = string(jsonBytes)

	fmt.Println(jsonString)
}
