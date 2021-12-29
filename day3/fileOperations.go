package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.Create("Data.txt")

	bdata := []byte("Hi all welcome to Golang")

	// now it is in write mode
	f.Write(bdata)
	fmt.Println("Data written")

	
}