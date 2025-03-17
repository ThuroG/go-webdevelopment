package models

import (
	"database/sql"
	"fmt"

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

//Open will open a SQL connection to the database
// Callers of Open need to ensure to close the connection eventually with db.Close() method
func Open(config PostgresConfig) (*sql.DB, error) {
	db, err := sql.Open("pgx", config.String())
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func DefaultPostgresConfig() PostgresConfig {
	return PostgresConfig{
	Host: "localhost",
	Port: "5432",
	User: "thuro",
	Password: "junglebook",
	Database: "webapp",
	SSLmode: "disable",
	}
}


func (cfg PostgresConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
	cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SSLmode)
}