package resolver

import (
	"github.com/OscarYuen/go-graphql-starter/model"
	"github.com/neelance/graphql-go"
)

type usersEdgeResolver struct {
	cursor graphql.ID
	model  *model.User
}

func (r *usersEdgeResolver) Cursor() graphql.ID {
	return r.cursor
}

func (r *usersEdgeResolver) Node() *userResolver {
	return &userResolver{u: r.model}
}
