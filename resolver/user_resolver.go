package resolver

import (
	"../model"
	"../service"
	"github.com/neelance/graphql-go"
	"strconv"
	"time"
)

type Resolver struct{}

type userResolver struct {
	u *model.User
}

func (r *userResolver) ID() graphql.ID {
	return graphql.ID(strconv.FormatInt(r.u.ID, 10))
}

func (r *userResolver) Email() *string {
	return &r.u.Email
}

func (r *userResolver) Password() *string {
	maskedPassword := "********"
	return &maskedPassword
}

func (r *userResolver) IPAddress() *string {
	return &r.u.IPAddress
}

func (r *userResolver) CreatedAt() (*graphql.Time, error) {
	if r.u.CreatedAt == "" {
		return nil, nil
	}

	t, err := time.Parse(time.RFC3339, r.u.CreatedAt)
	return &graphql.Time{Time: t}, err
}

type usersConnectionResolver struct {
	users      []*model.User
	totalCount int
	from       int
	to         int
}

func (r *usersConnectionResolver) TotalCount() int32 {
	return int32(r.totalCount)
}

func (r *usersConnectionResolver) Edges() *[]*usersEdgeResolver {
	l := make([]*usersEdgeResolver, r.to-r.from+1)
	for i := range l {
		l[i] = &usersEdgeResolver{
			cursor: service.EncodeCursor(r.from + i),
			model:  r.users[i],
		}
	}
	return &l
}

func (r *usersConnectionResolver) PageInfo() *pageInfoResolver {
	return &pageInfoResolver{
		startCursor: service.EncodeCursor(r.from),
		endCursor:   service.EncodeCursor(r.to),
		hasNextPage: r.to < r.totalCount,
	}
}

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

type pageInfoResolver struct {
	startCursor graphql.ID
	endCursor   graphql.ID
	hasNextPage bool
}

func (r *pageInfoResolver) StartCursor() *graphql.ID {
	return &r.startCursor
}

func (r *pageInfoResolver) EndCursor() *graphql.ID {
	return &r.endCursor
}

func (r *pageInfoResolver) HasNextPage() bool {
	return r.hasNextPage
}
