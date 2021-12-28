package main

import "fmt"

type employee struct {
	firstName string
	lastName string
	age int
}

func main() {
	x := employee{firstName: "R", lastName: "C", age: 25,}
	fmt.Printf("%v, %T", x, x)	
	
	rect := new(Rectangle) // rect1 is pointer to an instance of Rectangle
	rect.length = 10
	rect.width = 10
	rect.color = "red"
	fmt.Println(rect)

	var rect2 = &Rectangle{}
	rect2.length = 18
	rect2.width = 4
	fmt.Println(rect2)

	var rect3 = &Rectangle{}
	(*rect3).width = 2
	(*rect3).length = 7
	(*rect3).color = "Red"
	fmt.Println(rect3)

	s := Shape{length:4,
		rect: Rectangle{12,4,"Teal"}}
	fmt.Println(s)
	fmt.Println(s.length)
	fmt.Println(s.rect.color)
}

type Rectangle struct {
	length int
	width int
	color string
}

type Shape struct {
	length int
	rect Rectangle
}