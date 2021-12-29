package main

import (
	"fmt"
	"time"
)

func main() {
	ptime := time.Now()
	fmt.Println(ptime)

	p := fmt.Println
	p(time.Now().Year())
	p(time.Now().Month())
	p(time.Now().Hour())
	p(time.Now().Zone())

	wday := time.Now().Weekday()
	fmt.Println(wday)
	fmt.Println(int(wday))
}