package context

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"time"
)

func OpenDB(config *Config) (*sqlx.DB, error) {
	log.Println("Database is connecting... ")
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName))

	if err != nil {
		panic(err.Error())
	}

	if err = db.Ping(); err != nil {
		log.Println("Retry database connection in 5 seconds... ")
		time.Sleep(time.Duration(5) * time.Second)
		return OpenDB(config)
	}
	log.Println("Database is connected ")
	return db, nil
}
