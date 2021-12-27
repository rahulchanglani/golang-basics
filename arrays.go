package main

import "fmt"

func main() {
	arr := []int{2,4,5,6,7}
	for i,data := range arr{
		fmt.Println(i, " ", data)
	}
}