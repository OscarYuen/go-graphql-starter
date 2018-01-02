package service

import (
	"../model"
	"github.com/jmoiron/sqlx"
	"sync"
)

var UserService *userService
var once sync.Once

type userService struct {
	DB *sqlx.DB
}

func NewUserService(db *sqlx.DB) *userService {
	once.Do(func() {
		UserService = &userService{DB: db}
	})
	return UserService
}

func (u *userService) FindByEmail(email string) (*model.User, error) {
	user := &model.User{}
	row := u.DB.QueryRowx("SELECT * FROM users WHERE email=?", email)
	err := row.StructScan(user)
	return user, err
}

func (u *userService) CreateUser(user *model.User) (*model.User, error) {
	userSQL := `INSERT INTO users (email, password, ip_address) VALUES (:email, :password, :ip_address")`
	user.HashedPassword()
	_, err  := u.DB.NamedQuery(userSQL, user)
	if err != nil {
		return nil, err
	}
	return user,err
}


