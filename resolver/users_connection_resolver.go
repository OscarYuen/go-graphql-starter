package resolver

import (
	"github.com/OscarYuen/go-graphql-starter/model"
	"github.com/OscarYuen/go-graphql-starter/service"
)

type usersConnectionResolver struct {
	users      []*model.User
	totalCount int
	from       *string
	to         *string
}

func (r *usersConnectionResolver) TotalCount() int32 {
	return int32(r.totalCount)
}

func (r *usersConnectionResolver) Edges() *[]*usersEdgeResolver {
	l := make([]*usersEdgeResolver, len(r.users))
	for i := range l {
		l[i] = &usersEdgeResolver{
			cursor: service.EncodeCursor(&(r.users[i].ID)),
			model:  r.users[i],
		}
	}
	return &l
}

func (r *usersConnectionResolver) PageInfo() *pageInfoResolver {
	return &pageInfoResolver{
		startCursor: service.EncodeCursor(r.from),
		endCursor:   service.EncodeCursor(r.to),
		hasNextPage: false,
	}
}
