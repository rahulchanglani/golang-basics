package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOOS) // os details
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.GOROOT())
	fmt.Println(runtime.Compiler)
	fmt.Println(runtime.NumCPU()) // gives number of cores (cpu)
	fmt.Println(runtime.NumGoroutine())  // gives number of go routines
	fmt.Println(runtime.GOROOT())
	fmt.Println(runtime.Compiler)
}