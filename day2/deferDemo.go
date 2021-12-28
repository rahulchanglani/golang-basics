package main

import "fmt"

func main() {
	defer first()
	second()
}

func first() {
	fmt.Println("i will execute first")
}

func second() {
	fmt.Println("i will execute second")
}