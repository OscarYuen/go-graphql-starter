package schema

import (
	//"../conf"
	//"../model"
	"../service"
	"golang.org/x/net/context"
)

var userQuery = `
	findUserByEmail(email: String!): User
`

func (r *Resolver) FindUserByEmail(ctx context.Context,args struct {
	Email string
}) (*userResolver, error) {
	user, err := service.UserService.FindByEmail(ctx, args.Email)
	if err != nil {
		return nil, err
	}
	return &userResolver{user}, nil
}
