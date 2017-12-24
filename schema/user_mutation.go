package schema

import (
	"../conf"
	"../model"
	"../repository"
	"golang.org/x/net/context"
)

var userMutation = `
	createUser(email: String!, password: String!): User
`

func (r *Resolver) CreateUser(ctx context.Context,args *struct {
	Email    string
	Password string
}) (*userResolver, error) {
	user := &model.User{
		Email:    args.Email,
		Password: args.Password,
	}
	rb := &repository.UserRepository{repository.BaseRepository{DB:conf.DB}}
	result := rb.CreateUser(user)
	if  err := result.Error; err != nil {
		return nil, err
	}
	return &userResolver{user}, nil
}
