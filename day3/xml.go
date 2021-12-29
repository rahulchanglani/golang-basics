package main

import (
	"encoding/xml"
	"fmt"
)

var myxml = []byte(
	`<comp>
	<emp>
	<fname>Rohit</fname>
	<lname>C</lname>
	<id>123</id>
	<loc>Mys</loc>
	<emob>989898989</emob>
	<emob>989898980</emob>
	</emp>
	<emp>
	<fname>Sagar</fname>
	<lname>C</lname>
	<id>345</id>
	<loc>Ban</loc>
	<emob>989898989</emob>
	<emob>989898980</emob>
	</emp>
	</comp>`)

// create structure
type mycomp struct {
	Emp []employee `xml:"emp"`
}

type employee struct {
	// all variables should be public first letter capital
	Ename string `xml:"fname"`
	Eid   int    `xml:"id"`
	Eloc  string `xml:"loc"`
	Emob []string `xml:"emob"`
}

func main() {
	var empObj mycomp
	xml.Unmarshal(myxml, &empObj)
	fmt.Println(empObj)
}