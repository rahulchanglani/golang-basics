package main

import "fmt"

var i int = 27
func main() {
	fmt.Println(i)
	// i:=44 // reassigning is allowed if two same variables are of different scope
	var i int = 66
	fmt.Println(i)

	var num int = 445
	fmt.Println(num)
	num, d := 32, 20
	// on creating new variable, able to override
	fmt.Println(num, d)
}