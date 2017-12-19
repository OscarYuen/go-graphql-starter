package schema

import (
	"../conf"
	"../model"
	"../repository"
	"fmt"
)

var userQuery = `
	findUserByEmail(email: String!): User
`

func (r *Resolver) FindUserByEmail(args struct {
	Email string
}) (*userResolver, error) {
	user := &model.User{}
	result := repository.BaseRepository{DB: conf.DB}.Read(user, "9")
	if err := result.Error; err != nil {
		return nil, err
	}
	fmt.Printf("%v", result)
	return &userResolver{user}, nil
}
