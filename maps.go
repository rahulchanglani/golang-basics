package main

import "fmt"

func main() {
	myMap := map[int]string{
		101: "first",
		102: "second",
		103: "third",
	}

	fmt.Println(myMap[102])
	fmt.Printf("%T,%v", myMap, myMap)

	myMap[104]="forth"
	fmt.Println(myMap[104])
}
