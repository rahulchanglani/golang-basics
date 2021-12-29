package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt" // use cmd :::::: go get golang.org/x/crypto/bcrypt
)

func main() {
	pass := "Admin"
	newPass, _ := bcrypt.GenerateFromPassword([]byte(pass), 4)
	fmt.Println(newPass)

	// cost, _ := bcrypt.Cost(newPass)
	// fmt.Println(cost)

	// decryption
	err := bcrypt.CompareHashAndPassword(newPass, []byte("Admmin"))
	if err != nil {
		fmt.Println("Invalid password")
	} else {
		fmt.Println("Welcome user. Password matched. Login now")
	}
}

/*
go mod init golang.org/x/crypto/bcrypt
go mod tidy
*/
