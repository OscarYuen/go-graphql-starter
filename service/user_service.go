package service

import (
	"database/sql"
	"errors"
	"github.com/OscarYuen/go-graphql-starter/context"
	"github.com/OscarYuen/go-graphql-starter/model"
	"github.com/jmoiron/sqlx"
	"github.com/op/go-logging"
	"github.com/rs/xid"
	"log"
)

const (
	defaultListFetchSize = 10
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

	userSQL := `SELECT * FROM users WHERE email = $1`
	udb := u.db.Unsafe()
	row := udb.QueryRowx(userSQL, email)
	err := row.StructScan(user)
	if err == sql.ErrNoRows {
		return user, nil
	}
	if err != nil {
		u.log.Errorf("Error in retrieving user : %v", err)
		return nil, err
	}

	roles, err := u.roleService.FindByUserId(&user.ID)
	if err != nil {
		u.log.Errorf("Error in retrieving roles : %v", err)
		return nil, err
	}
	user.Roles = roles
	return user, nil
}

func (u *UserService) CreateUser(user *model.User) (*model.User, error) {
	userId := xid.New()
	user.ID = userId.String()
	userSQL := `INSERT INTO users (id, email, password, ip_address) VALUES (:id, :email, :password, :ip_address)`
	user.HashedPassword()
	_, err := u.db.NamedExec(userSQL, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserService) List(first *int32, after *string) ([]*model.User, error) {
	users := make([]*model.User, 0)
	var fetchSize int32
	if first == nil {
		fetchSize = defaultListFetchSize
	} else {
		fetchSize = *first
	}

	if after != nil {
		userSQL := `SELECT * FROM users WHERE created_at < (SELECT created_at FROM users WHERE id = $1) ORDER BY created_at DESC LIMIT $2;`
		decodedIndex, _ := DecodeCursor(after)
		err := u.db.Select(&users, userSQL, decodedIndex, fetchSize)
		if err != nil {
			return nil, err
		}
		return users, nil
	}
	userSQL := `SELECT * FROM users ORDER BY created_at DESC LIMIT $1;`
	err := u.db.Select(&users, userSQL, fetchSize)
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
		return nil, errors.New(context.UnauthorizedAccess)
	}
	if result := user.ComparePassword(userCredentials.Password); !result {
		return nil, errors.New(context.UnauthorizedAccess)
	}
	return user, nil
}
