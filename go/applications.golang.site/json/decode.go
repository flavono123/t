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

type DecodeMember struct {
	Name string
	Age  int
}

func main() {
	jsonBytes, _ := json.Marshal(Member{"Tim", 1, true})

	var mem DecodeMember
	err := json.Unmarshal(jsonBytes, &mem)
	if err != nil {
		panic(err)
	}

	fmt.Println(mem.Name, mem.Age) //, mem.Active)
}
