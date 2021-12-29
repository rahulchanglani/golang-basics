package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

func main() {
	info := "My name is Sam"
	fmt.Println(info)
	fmt.Println(strings.Compare(info, "My name is Sam"))
	fmt.Println(strings.Contains(info, "Sam"))
	fmt.Println(strings.Count(info, "a"))
	fmt.Println(strings.Split(info, "a"))
	fmt.Println(strings.Split("hi,hello,how,are,you?", ","))
	fmt.Println(strings.Split(info, "S"))
	fmt.Println(len(info))
	fmt.Printf("%T\n", strings.Split(info, "S"))
	fmt.Println(strings.Index(info, "i")) // index starts from 0
	fmt.Println(strings.Index(info, " "))

	////
	words := [...]string{"Seven", "even", "Maven", "Amen", "eleven"}
	for _, word := range words {
		found, err := regexp.MatchString(".even", word)
		if err != nil {
			log.Fatal(err)
		}
		if found {
			fmt.Printf("%s matches\n", word)
		} else {
			fmt.Printf("%s does not match\n", word)
		}
	}
}

/*
based on ascii lexical order


Return 0, if str1 == str2
Return 1, if str1 > str2
Return -1, if str1 < str2
*/

/*
str = "my name is hrushi"

with space: [my name is hrushi]
with blank: [m y  n a m e  i s  h r u s h i]
with letter i: [my name  s hrush ]
with space index 2: is
type of split with blank is: []string
len of split with blank is: 17
len of split with space is: 4
len of split with blank and index 1 is: 1
len of split with space and index 1 is: 4

*/