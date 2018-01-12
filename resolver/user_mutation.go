package resolver

import (
	"github.com/OscarYuen/go-graphql-starter/model"
	"github.com/OscarYuen/go-graphql-starter/service"
	"golang.org/x/net/context"
)

func (r *Resolver) CreateUser(ctx context.Context, args *struct {
	Email    string
	Password string
}) (*userResolver, error) {
	user := &model.User{
		Email:     args.Email,
		Password:  args.Password,
		IPAddress: *ctx.Value("requester_ip").(*string),
	}
	user, err := ctx.Value("userService").(*service.UserService).CreateUser(user)
	if err != nil {
		return nil, err
	}
	return &userResolver{user}, nil
}
