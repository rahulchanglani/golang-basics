package main

import "fmt"

// more than one defer, functions are there it will execute in LIFO

func main() {
	fmt.Println(div(12,3))
	fmt.Println(div(12,0))
	fmt.Println(div(12,4))
}

func div(n1, n2 int) int {
	defer func() {
		fmt.Println(recover())
	}()
	solution := n1 / n2
	return solution
}

// func demPanic() {
// 	defer func() {
// 		fmt.Println(recover())
// 	}()
// }