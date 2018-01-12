package resolver

import (
	"../config"
	"../model"
	"../service"
	"errors"
	"golang.org/x/net/context"
)

func (r *Resolver) User(ctx context.Context, args struct {
	Email string
}) (*userResolver, error) {
	user, err := ctx.Value("userService").(*service.UserService).FindByEmail(args.Email)
	if err != nil {
		return nil, err
	}
	return &userResolver{user}, nil
}

func (r *Resolver) Users(ctx context.Context, args struct {
	First *int32
	After *string
}) (*usersConnectionResolver, error) {
	if isAuthorized := ctx.Value("is_authorized").(bool); !isAuthorized {
		return nil, errors.New(config.CredentialsError)
	}
	notificationHub := ctx.Value("notificationHub").(*model.NotificationHub)
	noti := &model.Notification{From: 1, To: 9, Message: "1234222"}
	notificationHub.BroadcastMessage(noti)

	first := int(*args.First)
	users, err := ctx.Value("userService").(*service.UserService).List(&first, args.After)
	count, err := ctx.Value("userService").(*service.UserService).Count()
	if err != nil {
		return nil, err
	}
	return &usersConnectionResolver{users: users, totalCount: count, from: int(users[0].ID), to: int((users[len(users)-1]).ID)}, nil
}
