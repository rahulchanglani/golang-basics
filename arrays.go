package main

import "fmt"

func main() {
	arr := []int{2, 4, 5, 6, 7}
	for _, data := range arr {
		fmt.Println(">> ", data)
	}

	// string array
	name := []string{"Om", "RS"}
	for _, n := range name {
		fmt.Println(" ", n)
	}

	var intArr [3]int
	intArr[0] = 11
	intArr[1] = 111
	intArr[2] = 1111
	fmt.Println(intArr)
}
