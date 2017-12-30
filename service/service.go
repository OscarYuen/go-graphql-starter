package service

import "github.com/jmoiron/sqlx"

type Service interface {
	FindById(id string) *sqlx.DB
}

type service struct {
	//DB *sqlx.DB
}
