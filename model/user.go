package model

import (
	"time"
)

type User struct {
	ID        int64  `gorm:"primary_key"`
	Email     string `gorm:"size:255;not null;unique;index"`
	Password  string `gorm:"size:16;not null"`
	CreatedAt *time.Time
}
