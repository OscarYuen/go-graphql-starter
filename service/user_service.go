package service

import (
	"../model"
	"github.com/jmoiron/sqlx"
	"sync"
	"encoding/base64"
	"strconv"
	"strings"
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
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userService) CreateUser(user *model.User) (*model.User, error) {
	userSQL := `INSERT INTO users (email, password, ip_address) VALUES (:email, :password, :ip_address)`
	user.HashedPassword()
	_, err := u.DB.NamedExec(userSQL, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userService) List(first *int, after *string) ([]*model.User, error) {
	users := []*model.User{}
	decodedIndex, _ := decodeCursor(after)
	userSQL := `SELECT * FROM users WHERE id > ? - 1 LIMIT ? `
	err := u.DB.Select(&users, userSQL, decodedIndex, first)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userService) Count() (int, error) {
	var count int
	userSQL := `SELECT count(*) FROM users`
	err := u.DB.Get(&count, userSQL)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func decodeCursor(after *string) (*int, error){
	decodedValue := 0
	if after != nil {
		b, err := base64.StdEncoding.DecodeString(string(*after))
		if err != nil {
			return nil, err
		}
		i, err := strconv.Atoi(strings.TrimPrefix(string(b), "cursor"))
		if err != nil {
			return nil, err
		}
		decodedValue = i
	}
	return &decodedValue,nil
}


