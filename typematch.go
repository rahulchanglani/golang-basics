package main

import "fmt"

func main() {
	var num1 int16 = 20
	var num2 int32 = 33
	fmt.Println(num1 + int16(num2))
}