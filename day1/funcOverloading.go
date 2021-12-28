package main

import "fmt"

func main() {
	add(12, 3)
	m, d := multi(12, 3)
	fmt.Println(m, d)
}

func add(num1,num2 int) {
	fmt.Println("sum is", num1+num2)
}

func multi(n1 int, n2 int) (int, int) {
	return n1 * n2, n1/n2
}