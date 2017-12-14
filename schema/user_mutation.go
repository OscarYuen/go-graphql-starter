package schema

import (
	"../model"
)

func (r *Resolver) CreateUser(args *struct {
	Email    string
	Password string
}) (*userResolver, error) {
	user := &model.User{
		Email:    args.Email,
		Password: args.Password,
	}
	if err := DB.Create(user).Error; err != nil {
		return nil, err
	}
	return &userResolver{user}, nil
}
