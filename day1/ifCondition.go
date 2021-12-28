package main

import "fmt"

func main() {
	age:=20
	if age < 15 {
		fmt.Println("if")
	} else if age == 15 {
		fmt.Println("else if")
	} else {
		fmt.Println("else")
	}

	var n1,n2,n3 int
	fmt.Scanln(&n1, &n2, &n3)
		if n1>n2 && n1>n3 {
			fmt.Println("n1 is largest")
		} else if n2>n3 && n2>n1 {
			fmt.Println("n2 is largest")
		} else {
			fmt.Println("n3 is largest")
		}
}