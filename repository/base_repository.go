package repository

import (
	"github.com/jinzhu/gorm"
)

type BaseRepository struct {
	DB *gorm.DB
}

func (b BaseRepository) Read(pointerOfEntities interface{}, id string) *gorm.DB {
	return b.DB.Where("id = ?", id).First(pointerOfEntities)
}

func (b BaseRepository) ReadAll(pointerOfEntities interface{}) *gorm.DB {
	return b.DB.Find(pointerOfEntities)
}

func (b BaseRepository) Create(pointerOfEntities interface{}) *gorm.DB {
	return b.DB.Create(pointerOfEntities)
}
