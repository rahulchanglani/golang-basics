package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// check file is there or not - if file exists then append or create new file
	fp, _ := os.OpenFile("Data.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	fp.WriteString("\n we are working on Golang")
	fmt.Println("done")
	fp.Close()   // to close file


	///// read
	content, err := os.ReadFile("Data.txt")
	if (err != nil) {
		fmt.Println(err)
	}
	fmt.Println(string(content))
	fp.Close()


	//
	d, _ := ioutil.ReadFile("Datanew.txt")
	fmt.Println(string(d))
}