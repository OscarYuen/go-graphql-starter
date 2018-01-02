package schema

import (
	//"../conf"
	"../model"
	"../service"
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
		IPAddress: ctx.Value("requester_ip").(string),
	}
	user, err := service.UserService.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return &userResolver{user}, nil
}
