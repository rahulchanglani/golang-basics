package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.Create("Data.txt") // checks and creates if not present

	bdata := []byte("Hi all welcome to Golang")

	// now it is in write mode
	f.Write(bdata)
	fmt.Println("Data written")

	/// another method
	f.WriteString("\n Acc")
	fmt.Println("Data updated")
}