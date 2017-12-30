package service

import (
	"../model"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
)

var UserService = &userService{}

type userService struct {
	//DB *sqlx.DB
}

func (u *userService) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	user := &model.User{}
	db := ctx.Value("db").(*sqlx.DB)
	row := db.QueryRowx("SELECT * FROM users WHERE email=?", email)
	err := row.StructScan(user)
	return user, err
}

//func (u *userService) CreateUser(user *model.User) *gorm.DB {
//	user.HashedPassword()
//	return u.DB.Create(user)
//}


