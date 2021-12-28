package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	wg.Add(2) // add and give no of goroutines, wait grp will wait for 2 goroutines
	go sayHello()
	go sayBye()
	fmt.Println("Waiting inside main")
	wg.Wait() // wait till go routine acknowledge as Done
	fmt.Println("last line in main")
}

func sayHello() {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello")
		time.Sleep(1000 * time.Millisecond)
	}
	wg.Done()
}

func sayBye() {
	for i := 0; i < 5; i++ {
		fmt.Println("bye")
		time.Sleep(1000 * time.Millisecond)
	}
	wg.Done()
}
