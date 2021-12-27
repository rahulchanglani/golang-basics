package main

import "fmt"

const PI = 3.14
const ID int = 89898

var (
	num1 int16 = 20
	num2 int32 = 33
)

func main() {
	fmt.Println(num1 + int16(num2))
	var n1, n2, n3 int = 12, 13, 14
	fmt.Println(n1 + n2 + n3)

	loc, sal := "BDC", 222
	fmt.Println(loc, sal)
	fmt.Println(PI, ID)
}
