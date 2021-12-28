package main

import "fmt"

func variadic(a string, params ...int) {
	fmt.Printf("%T, %v", params, params)
}

func main() {
	variadic("hi", 12, 13)
	sum(1,2)
	sum(3,4,5,6)
	nums := []int{2,4,5,6}
	// sum(nums...)
	sum(nums[1:3]...)
}

func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}
