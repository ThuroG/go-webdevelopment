package main

import (
	"database/sql"
	"fmt"
	"go-webdevelopment/models"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type PostgresConfig struct {
	Host string
	Port string
	User string
	Password string
	Database string
	SSLmode string
}

func Main() {
	// Define the PostgresConfig struct

	cfg := PostgresConfig{
		Host: "localhost",
		Port: "5432",
		User: "thuro",
		Password: "junglebook",
		Database: "webapp",
		SSLmode: "disable",
	}
	// Use pgx in order to connect to Postgresql
	db, err := sql.Open("pgx", cfg.String())
	if err != nil {
		panic(err)
	}
	defer db.Close() //Close connection if err occurred
	err = db.Ping() // throws an error if not pingable
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected")

	us := models.UserService{
		DB: db,
	  }
	  user, err := us.Create("bob@bob.com", "bob123")
	  if err != nil {
		panic(err)
	  }
	  fmt.Println(user)
}