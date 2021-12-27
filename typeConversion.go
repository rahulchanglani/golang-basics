package main

import "fmt"

func main() {
	var i int = 42
	fmt.Printf("%v, %T\n", i, i)

	var j string
	j = string(i)
	fmt.Printf("%v, %T\n", j, j)
}
