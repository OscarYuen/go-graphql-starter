package model

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

type User struct {
	ID        int64
	Email     string
	Password  string
	IPAddress string `db:"ip_address"`
	CreatedAt string `db:"created_at"`
}

func (user *User) HashedPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return err
	}
	user.Password = string(hash)
	return nil
}

func (user *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}


