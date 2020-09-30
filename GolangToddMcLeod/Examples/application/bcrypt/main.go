package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {

	// encrypt password
	s := "password123_strong"
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}

	// enter password and check against encypted
	var password string

	fmt.Println("Please enter your password:")
	fmt.Scanf("%s", &password)

	err = bcrypt.CompareHashAndPassword(bs, []byte(password))
	if err != nil {
		fmt.Println("Wrong Password")
	} else {
		fmt.Println("Login Successful!")
	}
}
