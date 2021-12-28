package main

import (
	"fmt"
	"time"
)

func main() {
	pt := time.Now()
	fmt.Println(pt)
	time.Sleep(5000*time.Millisecond)
	fmt.Println(pt.Date())
	fmt.Println(pt.Hour())
	fmt.Println(pt.Minute())
	//
	fmt.Println(pt.Local())
	//clock
	fmt.Println(pt.Clock())
}