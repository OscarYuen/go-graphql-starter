package schema

import (
	"github.com/OscarYuen/go-graphql-example/model"
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
	user.HashedPassword()
	if err := DB.Create(user).Error; err != nil {
		return nil, err
	}
	return &userResolver{user}, nil
}
