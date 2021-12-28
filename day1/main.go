package main

import "fmt"

var name = "Rahul"
var num1 int // this is possible in global and local scope

func main() {
	str2 := "abc" // Shorthand variable only possible in local scope
	fmt.Println(num1)
	fmt.Printf("%T", num1)
	fmt.Printf("%v", name)
	fmt.Printf("%q", str2)
	var num2 = 123.45
	fmt.Println(num2)
	var str5 string
	str5 = "xyz"
	fmt.Printf("%s", str5)
	var token string
	fmt.Println("\nEnter the token")
	fmt.Scanln(&token)
	fmt.Printf("Entered token is %s", token)
	fmt.Println("\nEnter two numbers")
	var n1, n2 int
	fmt.Scanln(&n1, &n2)
	fmt.Println("\nProduct of 2 nos is", (n1 * n2))
}