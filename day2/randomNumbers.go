package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Print(rand.Intn(10))
	fmt.Print()
	fmt.Print(rand.Float64())
}