package main

import "fmt"

func main() {
	hello()
	f1(6, 7)
	add(88, 8)
}

func hello() {
	fmt.Println("Hello")
}

func f1(a int, b int) {
	fmt.Println(a + b)
}

func add(a, b int) int {
	return a + b
}