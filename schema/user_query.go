package schema

import (
	"github.com/OscarYuen/go-graphql-example/model"
)

var userQuery = `
	findUserByEmail(email: String!): User
`

func (r *Resolver) FindUserByEmail(args struct {
	Email string
}) (*userResolver, error) {
	user := &model.User{}
	result := DB.Where("email = ?", args.Email).First(user)
	if err := result.Error; err != nil {
		return nil, err
	}
	return &userResolver{user}, nil
}
