package main

import "fmt"

func display(ans func(a, b int)int) {
	fmt.Println(ans(100,200))
}

func main() {
	func (n1 int, n2 int) {
		fmt.Println("The sum of 2 nos is",n1+n2)
	}(12,34)

	ans := func (a, b int) int {
		return a * b
	}
	display(ans)
}