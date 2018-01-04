package schema

import (
	"../service"
	"golang.org/x/net/context"
)

var userQuery = `
	user(email: String!): User
	users(first: Int,  after: String): UsersConnection!
`

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
	//user1 := &model.User{
	//	ID:        int64(123),
	//	Email:     "123",
	//	Password:  "123",
	//	IPAddress: "123",
	//}
	//user2 := &model.User{
	//	ID:        int64(124),
	//	Email:     "124",
	//	Password:  "124",
	//	IPAddress: "124",
	//}
	//users := []*model.User{user1, user2}
	first := int(*args.First)
	users,_ := ctx.Value("userService").(*service.UserService).List(&first, args.After)
	count,_ := ctx.Value("userService").(*service.UserService).Count()
	return &usersConnectionResolver{users: users,totalCount:count, from: int(users[0].ID), to: int((users[len(users)-1]).ID)}, nil
}
