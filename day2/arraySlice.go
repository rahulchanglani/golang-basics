package main

import "fmt"

func main() {
	a := []int{1,2,3,4,5}

	fmt.Println(a[:])  // slice all elements
	// sum(nums...)
	fmt.Println(a[0:4])  // 0 to 3, excluding 4
	fmt.Println(a[4:])  // 4 onwards
	fmt.Println(a[:6])
}