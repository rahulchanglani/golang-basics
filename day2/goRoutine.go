package main

import (
	"fmt"
	"time"
)

func main() {
	go sayhello("o")
	sayhello("hi")
}

func sayhello(s string) {
	for i:=0;i<5;i++{
		time.Sleep(100*time.Millisecond)
		fmt.Println(s)
	}
}