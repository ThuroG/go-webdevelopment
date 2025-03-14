package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main(){
	for i, arg := range os.Args{
		fmt.Println(i, arg)
	}
	switch os.Args[1]{
	case "hash":
		hash(os.Args[2])
	case "compare":
		compare(os.Args[2], os.Args[3])
	default:
		fmt.Println("Usage: bcrypt <hash|compare> [args]")
	}
}

func hash(password string) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		fmt.Printf("Password %v has a Hasherror: %v\n ", password, err)
		return
	}
	hash := string(hashedBytes)
	fmt.Printf("Password: %s\nHashed: %s\n", password, hash)
}

func compare(password, hash string) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		fmt.Printf("Password is invalid %v\n", password)
		return
	}
	fmt.Printf("Password is correct")
}