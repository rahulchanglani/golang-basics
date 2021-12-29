package main

import (
	"fmt"
	"encoding/json"
)

type Emp struct {
	Id int
	Name string
}

func main() {
	text := "[{\"id\":100,\"name\":\"hello\"}]"
	jsonData := []byte(text)
	var empObj []Emp
	json.Unmarshal(jsonData, &empObj)

	for e := range empObj {
		fmt.Printf("Id = %v, Name = %v", empObj[e].Id, empObj[e].Name)
		fmt.Println()
	}
}