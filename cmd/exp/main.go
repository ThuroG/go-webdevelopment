package main

import (
	"errors"
	"fmt"
	"html/template"
	"os"
	"embed"


)

//go:embed *
var FS embed.FS

var ErrNotFound = errors.New("not found")

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

// SOLVED: HAS TO BE RUN IN CMD/EXP FOLDER otherwise it will not work
func main() {
	// Section 7 use ParseFS to call it from anywhere
	t := template.Must(template.ParseFS(FS, "hello.gohtml"))


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

	t.Execute(os.Stdout, user)


	// fmt.ErrorF will append all error messages so that it is easier to debug
	// use err1 because err already in use before
	err1:= CreateOrg()
	if err1 != nil {
		fmt.Println(err1)
	}
	
	// Section 6 - Exercise: Use Errors.Is to detect error kind (see global variable)
	err := B()
	if errors.Is(err, ErrNotFound) {
		fmt.Println("This error has been indeed not found (yet)")
	} else {
		fmt.Println(err)
	}
  }
  

  
  func A() error {
	  return ErrNotFound
  }
  
  func B() error {
	  err := A()
	if err != nil {
		return fmt.Errorf("b: %w", err)
	}
	return nil
  }