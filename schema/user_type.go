package schema

import (
	"github.com/OscarYuen/go-graphql-example/model"
	"strconv"
)

var userSchema = `
	type User {
		id: String!
		email: String!
		password: String!
		createdAt: String!
	}`

type userResolver struct {
	u *model.User
}

func (r *userResolver) ID() string {
	return strconv.FormatInt(r.u.ID, 10)
}

func (r *userResolver) Email() string {
	return r.u.Email
}

func (r *userResolver) Password() string {
	return r.u.Password
}

func (r *userResolver) CreatedAt() string {
	return r.u.CreatedAt.String()
}
