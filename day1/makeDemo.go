package main

import "fmt"

func main() {
	var spiritual = make(map[string]string)
	spiritual["RS"] = "Beas"
	spiritual["Om"] = "Coimbature"
	fmt.Println(spiritual)

	//array
	myVar := make([]string,2,5)
	fmt.Println(len(myVar))
	fmt.Println(cap(myVar))
	myVar[0]="AMD"
	myVar[1]="TN"
	fmt.Println(myVar)

	//
	myVar = append(myVar, "22")
	myVar = append(myVar, "25")
	myVar = append(myVar, "27")
	fmt.Println(myVar)
	fmt.Println(cap(myVar))
}