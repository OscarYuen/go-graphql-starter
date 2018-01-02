package model

import (
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type User struct {
	ID        int64
	Email     string
	Password  string
	IPAddress string `db:"ip_address"`
	CreatedAt *time.Time `db:"created_at"`
}

func (user *User) HashedPassword() {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
		return
	}
	user.Password = string(hash)
}

