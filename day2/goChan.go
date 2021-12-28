package main

import "fmt"

func main() {
	done := make(chan bool)
	go display(done) // goroutine
	<- done // either store or print
}

func display(done chan bool) {
	fmt.Println("display goroutine")
	done <- true
}