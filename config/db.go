package config

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

func OpenDB(path string) (*sqlx.DB, error) {
	log.Println("Database is initializing... ")
	db, err := sqlx.Connect("postgres", "host=localhost port=5432 user=postgres password=123456 dbname=postgres sslmode=disable")
	log.Println("Database is initialized ")
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
