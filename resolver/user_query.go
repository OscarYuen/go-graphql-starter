package resolver

import (
	"errors"
	"github.com/OscarYuen/go-graphql-starter/config"
	"github.com/OscarYuen/go-graphql-starter/loader"
	"github.com/OscarYuen/go-graphql-starter/service"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
)

func (r *Resolver) User(ctx context.Context, args struct {
	Email string
}) (*userResolver, error) {
	//Without using dataloader:
	//user, err := ctx.Value("userService").(*service.UserService).FindByEmail(args.Email)
	userId := ctx.Value("user_id").(*int64)
	user, err := loader.LoadUser(ctx, args.Email)
	if err != nil {
		ctx.Value("logger").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("logger").(*logging.Logger).Debugf("Retrieved user by user_id[%d] : %v", *userId, *user)
	return &userResolver{user}, nil
}

func (r *Resolver) Users(ctx context.Context, args struct {
	First *int32
	After *string
}) (*usersConnectionResolver, error) {
	if isAuthorized := ctx.Value("is_authorized").(bool); !isAuthorized {
		return nil, errors.New(config.CredentialsError)
	}
	userId := ctx.Value("user_id").(*int64)

	first := int(*args.First)
	users, err := ctx.Value("userService").(*service.UserService).List(&first, args.After)
	count, err := ctx.Value("userService").(*service.UserService).Count()
	ctx.Value("logger").(*logging.Logger).Debugf("Retrieved users by user_id[%d] : %v", *userId, users)
	ctx.Value("logger").(*logging.Logger).Debugf("Retrieved total users count by user_id[%d] : %v", *userId, count)
	if err != nil {
		ctx.Value("logger").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	return &usersConnectionResolver{users: users, totalCount: count, from: int(users[0].ID), to: int((users[len(users)-1]).ID)}, nil
}
