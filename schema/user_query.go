package schema

import (
	"../model"
)

func (r *Resolver) FindUserByEmail(args struct {
	Email string
}) (*userResolver, error) {
	user := &model.User{
		Email: args.Email,
	}
	result := DB.Where("email = ?", args.Email).First(user)
	if err := result.Error; err != nil {
		return nil, err
	}
	return &userResolver{user}, nil
}
