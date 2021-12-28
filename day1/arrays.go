package main

import "fmt"

func main() {
	arr := [...]int{1,2,3,4,5}
	for _, data := range arr {
		fmt.Println(">> ", data)
	}
	array2 := arr
	array2[1] = 10
	fmt.Println(arr)
	fmt.Println(array2)

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
