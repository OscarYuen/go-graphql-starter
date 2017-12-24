package repository

import (
	"github.com/jinzhu/gorm"
	"../model"
)

type UserRepository struct {
	BaseRepository
}

func (u *UserRepository) FindByEmail(user *model.User, id string) *gorm.DB {
	return u.DB.Where("email = ?", id).First(user)
}

func (u *UserRepository) CreateUser(user *model.User) *gorm.DB {
	user.HashedPassword()
	return u.BaseRepository.Create(user)
}


