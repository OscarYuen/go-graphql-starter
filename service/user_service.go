package service

import (
	"database/sql"
	"errors"
	"github.com/OscarYuen/go-graphql-starter/config"
	"github.com/OscarYuen/go-graphql-starter/model"
	"github.com/jmoiron/sqlx"
	"github.com/op/go-logging"
)

const (
	defaultListFetchSize = 10
	defaultDecodedIndex  = 0
)

type UserService struct {
	db          *sqlx.DB
	roleService *RoleService
	log         *logging.Logger
}

func NewUserService(db *sqlx.DB, roleService *RoleService, log *logging.Logger) *UserService {
	return &UserService{db: db, roleService: roleService, log: log}
}

func (u *UserService) FindByEmail(email string) (*model.User, error) {
	user := &model.User{}

	userSQL := `SELECT user.*
	FROM users user
	WHERE user.email = ? `
	row := u.db.QueryRowx(userSQL, email)
	err := row.StructScan(user)
	if err == sql.ErrNoRows {
		return user, nil
	}
	if err != nil {
		u.log.Errorf("Error in retrieving user : %v", err)
		return nil, err
	}

	roles, err := u.roleService.FindByUserId(user.ID)
	if err != nil {
		u.log.Errorf("Error in retrieving roles : %v", err)
		return nil, err
	}
	user.Roles = roles
	return user, nil
}

func (u *UserService) CreateUser(user *model.User) (*model.User, error) {
	userSQL := `INSERT INTO users (email, password, ip_address) VALUES (:email, :password, :ip_address)`
	user.HashedPassword()
	_, err := u.db.NamedExec(userSQL, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserService) List(first *int, after *string) ([]*model.User, error) {
	users := make([]*model.User, 0)
	decodedIndex, _ := DecodeCursor(after)
	if first == nil {
		*first = defaultListFetchSize
	}
	userSQL := `SELECT * FROM users WHERE id > ? - 1 LIMIT ? `
	err := u.db.Select(&users, userSQL, decodedIndex, first)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserService) Count() (int, error) {
	var count int
	userSQL := `SELECT count(*) FROM users`
	err := u.db.Get(&count, userSQL)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (u *UserService) ComparePassword(userCredentials *model.UserCredentials) (*model.User, error) {
	user, err := u.FindByEmail(userCredentials.Email)
	if err != nil {
		return nil, errors.New(config.UnauthorizedAccess)
	}
	if result := user.ComparePassword(userCredentials.Password); !result {
		return nil, errors.New(config.UnauthorizedAccess)
	}
	return user, nil
}
