package main

import "fmt"

func main() {
	for i:=0;i<10;i++ {
		fmt.Println(i)
	}

	// while
	i := 50
	for i < 55 {
		fmt.Println(i)
		i++
	}

	////// break
	for j:=0;j<5;j++ {
		if j == 4 {
			break
		}
		fmt.Println(i)
	}

	// continue
	for j:=0;j<5;j++ {
		if j == 4 {
			continue
		}
		fmt.Println(j)
	}

}