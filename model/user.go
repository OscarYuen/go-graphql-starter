package model

import (
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type User struct {
	ID        int64  `gorm:"primary_key"`
	Email     string `gorm:"size:255;not null;unique;index"`
	Password  string `gorm:"size:16;not null"`
	CreatedAt *time.Time
}

func (user *User) HashedPassword() {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
		return
	}
	user.Password = string(hash)
}
