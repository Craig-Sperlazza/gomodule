package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// https://pkg.go.dev/golang.org/x/crypto/bcrypt#pkg-overview
// bcrypt package helps you encrypt and decrypt your passwords.


func main() {
	fmt.Println(2)
	s := `password123`
	// this function will encrypt your password
	// bcrypt.GenerateFromPassword(password []byte, cost int) ([]byte, error)
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(bs)

	// this function will let you know if a given password matches the hash
	// func CompareHashAndPassword(hashedPassword, password []byte) error

	// note you return after the bad password, so these both cant run...just for illustration

	badPassword := `eewewrwtr123`

	err2 := bcrypt.CompareHashAndPassword(bs, []byte(badPassword))
	if err2 != nil {
		fmt.Println("Error Bad Password")
		return
	}
	fmt.Println("Passwords Match")


	goodPass := `password123`
	err3 := bcrypt.CompareHashAndPassword(bs, []byte(goodPass))
	if err3 != nil {
		fmt.Println("Error Bad Password")
		return
	}
	fmt.Println("Passwords Match")
}