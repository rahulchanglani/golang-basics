package main

import "fmt"

type Address struct {
	pin int
	city string
}

type Student struct {
	id int
	name string
	age int
	sadd Address
}

func main() {
	a := Address{560001, "Bangalore"}
	fmt.Println(a)
	s := Student{123, "Rohit", 20, a}
	fmt.Println(s)
	//fmt.Println(s.id, s.sadd)
	sobj := Student{id:101, name:"Sam", age:23, sadd:Address{572106,"Mysore"}}
	fmt.Println(sobj.sadd.city)
}