package config

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"time"
)

func OpenDB(host string, port string, user string, password string, dbname string) (*sqlx.DB, error) {
	log.Println("Database is connecting... ")
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))

	if err != nil {
		panic(err.Error())
	}

	if err = db.Ping(); err != nil {
		log.Println("Retry database connection in 5 seconds... ")
		time.Sleep(time.Duration(5) * time.Second)
		return OpenDB(host, port, user, password, dbname)
	}
	log.Println("Database is connected ")
	return db, nil
}
