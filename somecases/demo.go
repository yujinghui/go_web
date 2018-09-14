package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	defer fmt.Println("finnally")
	var str = "hello world"
	fmt.Println(len(str))
	content := []byte(`{"hello":22, "world":2}`)
	var f interface{}
	err := json.Unmarshal(content, &f)
	if err != nil {
		panic(err)
	}
	fmt.Println(f)
}
