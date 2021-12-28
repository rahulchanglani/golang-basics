package main

import (
	"fmt"
)

func main() {
	// char datatype is not there in golang. use String with one character
	var str string
	
	switch str {
	case "a", "e", "i", "o", "u":
		fmt.Println("Vowel", str)
	default:
		fmt.Println("Consonant", str)
	}
}
