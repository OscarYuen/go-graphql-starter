package repository

import (
	"github.com/jinzhu/gorm"
)

type BaseRepository struct {
	DB *gorm.DB
}

type BaseRepositoryInterface interface {
	FindById(pointerOfEntities interface{}, id string) *gorm.DB
	ReadAll(pointerOfEntities interface{}) *gorm.DB
	Create(pointerOfEntities interface{}) *gorm.DB
}

func (b *BaseRepository) FindById(pointerOfEntities interface{}, id string) *gorm.DB {
	return b.DB.Where("id = ?", id).First(pointerOfEntities)
}

func (b *BaseRepository) ReadAll(pointerOfEntities interface{}) *gorm.DB {
	return b.DB.Find(pointerOfEntities)
}

func (b *BaseRepository) Create(pointerOfEntities interface{}) *gorm.DB {
	return b.DB.Create(pointerOfEntities)
}



