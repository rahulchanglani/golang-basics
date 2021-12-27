package main

import "fmt"

func main() {
	age:=20
	if age < 15 {
		fmt.Println("if")
	} else if age == 15 {
		fmt.Println("else if")
	} else {
		fmt.Println("else")
	}
}