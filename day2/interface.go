package main

import "fmt"

type vehicle interface {
	engine()
}

//create function
func vdetails(v vehicle) {
	fmt.Println(v)
	println(v)
	v.engine()
}

//create struct
type Maruti struct{
	model string
	name string
}

//override function then it will become vehicle
func (m Maruti) engine(){
	fmt.Println("Maruti has a auto engine")
}

func main(){
	// mobj := Maruti{"M10","Scross"}
	// mobj.engine()
	vdetails(Maruti{"S11","Baleno"})
	name:="Lata"
	println(name)
}