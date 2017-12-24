package schema

import (
	"../conf"
	"../model"
	"../repository"
)

var userQuery = `
	findUserByEmail(email: String!): User
`

func (r *Resolver) FindUserByEmail(args struct {
	Email string
}) (*userResolver, error) {
	user := &model.User{}
	rb := &repository.UserRepository{repository.BaseRepository{DB:conf.DB}}
	result := rb.FindByEmail(user, args.Email)
	if err := result.Error; err != nil {
		return nil, err
	}
	return &userResolver{user}, nil
}
