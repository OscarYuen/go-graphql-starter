package service

import (
	"database/sql"
	"github.com/OscarYuen/go-graphql-starter/model"
	"github.com/jmoiron/sqlx"
	"github.com/op/go-logging"
)

type RoleService struct {
	db  *sqlx.DB
	log *logging.Logger
}

func NewRoleService(db *sqlx.DB, log *logging.Logger) *RoleService {
	return &RoleService{db: db, log: log}
}

func (r *RoleService) FindByUserId(userId int64) ([]*model.Role, error) {
	roles := make([]*model.Role, 0)

	roleSQL := `SELECT role.*
	FROM roles role
	INNER JOIN rel_users_roles ur ON role.id = ur.role_id
	WHERE ur.user_id = ? `
	err := r.db.Select(&roles, roleSQL, userId)
	if err == sql.ErrNoRows {
		return roles, nil
	}
	if err != nil {
		return nil, err
	}
	return roles, nil
}
