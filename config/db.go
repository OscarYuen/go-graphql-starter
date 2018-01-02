package config

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func OpenDB(path string) (*sqlx.DB, error){
	log.Println("Start to open a new DB... ")
	db, err := sqlx.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
