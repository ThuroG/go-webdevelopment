package main

import (
	"os"
	"html/template"
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
}
