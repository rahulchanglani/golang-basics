package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 42
	fmt.Printf("%v, %T\n", i, i)

	var j string
	// j = string(i)    // gives * on converting to string
	j = strconv.Itoa(i) // I int to ascii
	fmt.Printf("%v, %T\n", j, j)

	var n1 int16 = 20
	fmt.Printf("%T, %v", int32(n1), int32(n1))
}
