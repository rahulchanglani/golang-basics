package main

import (
	"fmt"
)

func main() {
	a := makeRange(1,20)
	fmt.Println(a)
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	// fmt.Println(a)
	for i := range a{
		a[i] = min + i
	}
	return a
}