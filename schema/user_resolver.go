package schema

import (
	"../model"
	"strconv"
	"github.com/neelance/graphql-go"
	"time"
)

var userSchema = `
	scalar Time
	type User {
		id: ID!
		email: String
		password: String
		ipAddress: String
		createdAt: Time
	}`

type userResolver struct {
	u *model.User
}

func (r *userResolver) ID() graphql.ID {
	return graphql.ID(strconv.FormatInt(r.u.ID, 10))
}

func (r *userResolver) Email() *string {
	return &r.u.Email
}

func (r *userResolver) Password() *string {
	return &r.u.Password
}

func (r *userResolver) IPAddress() *string {
	return &r.u.IPAddress
}

func (r *userResolver) CreatedAt() (*graphql.Time, error) {
	if r.u.CreatedAt == "" {
		return nil, nil
	}

	t, err := time.Parse(time.RFC3339, r.u.CreatedAt)
	return &graphql.Time{Time: t}, err
}
