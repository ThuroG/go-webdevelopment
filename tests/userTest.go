package main

import (
	"errors"
	"fmt"
	"go-webdevelopment/models"
)

type PostgresConfig struct {
	Host string
	Port string
	User string
	Password string
	Database string
	SSLmode string
}


func (cfg PostgresConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SSLmode)
}

func Connect() error {
	//pretend we get an error
	return errors.New("failed to create connection")
}

func main() {
	// Define the PostgresConfig struct

	cfg := models.DefaultPostgresConfig()
	// Use pgx in order to connect to Postgresql
	db, err := models.Open(cfg)
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