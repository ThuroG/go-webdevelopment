package main

import (
	"os"
	"html/template"
	"errors"
	"fmt"
)

type User struct {
	Name string
	Bio string
	Age int
	Address Address
}

type Address struct {
	Street string
	ZipCode string
	City string
}

func Connect() error {
	//pretend we get an error
	return errors.New("failed to create connection")
}

func CreateUser() error {
	err := Connect()
	if err != nil {
		return fmt.Errorf("failed to create USER: %w", err)
	}
	return nil
}

func CreateOrg() error {
	err := CreateUser()
	if err != nil {
		return fmt.Errorf("failed to create org: %w", err)
	}
	return nil
}

//HAS TO BE RUN IN CMD/EXP FOLDER otherwise it will not work
func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}


	address := Address{
		Street: "123 Main St",
		ZipCode: "12345",
		City: "Springfield",
	}

	user := User{
		Name: "Arthur",
		Bio: `alert("Haha, you have been h4x0r3d!");`,
		Age: 123,
		Address: address,
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}

	// fmt.ErrorF will append all error messages so that it is easier to debug
	// use err1 because err already in use before
	err1:= CreateOrg()
	if err1 != nil {
		fmt.Println(err1)
	}

}