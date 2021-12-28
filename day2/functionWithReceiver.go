package main

import "fmt"

type lang []string

func (l lang) display() {
	for i, lang := range l{
		fmt.Println(i,lang)
	}
}

func main() {
	lang{"Python","Java","Kotlin"}.display() // array.functionname()
}