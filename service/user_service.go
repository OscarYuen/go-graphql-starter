package service

import (
	"../model"
	"github.com/jmoiron/sqlx"
	"sync"
)

const (
	defaultListFetchSize = 10
	defaultDecodedIndex  = 0
)

var (
	userServiceInstance *UserService
	userOnce            sync.Once
)

type UserService struct {
	DB *sqlx.DB
}

func NewUserService(db *sqlx.DB) *UserService {
	userOnce.Do(func() {
		userServiceInstance = &UserService{DB: db}
	})
	return userServiceInstance
}

func (u *UserService) FindByEmail(email string) (*model.User, error) {
	user := &model.User{}
	row := u.DB.QueryRowx("SELECT * FROM users WHERE email=?", email)
	err := row.StructScan(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserService) CreateUser(user *model.User) (*model.User, error) {
	userSQL := `INSERT INTO users (email, password, ip_address) VALUES (:email, :password, :ip_address)`
	user.HashedPassword()
	_, err := u.DB.NamedExec(userSQL, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserService) List(first *int, after *string) ([]*model.User, error) {
	users := []*model.User{}
	decodedIndex, _ := DecodeCursor(after)
	if first == nil {
		*first = defaultListFetchSize
	}
	userSQL := `SELECT * FROM users WHERE id > ? - 1 LIMIT ? `
	err := u.DB.Select(&users, userSQL, decodedIndex, first)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserService) Count() (int, error) {
	var count int
	userSQL := `SELECT count(*) FROM users`
	err := u.DB.Get(&count, userSQL)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (u *UserService) ComparePassword(userCredentials *model.UserCredentials) (bool, *model.User) {
	user, err := u.FindByEmail(userCredentials.Email)
	if err != nil {
		return false, nil
	}
	return user.ComparePassword(userCredentials.Password), user
}
