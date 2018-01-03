package schema

import (
	"../model"
	"encoding/base64"
	"fmt"
	"github.com/neelance/graphql-go"
	"strconv"
	"time"
)

var userSchema = `
	scalar Time
	type User {
		id: ID!
		email: String
		password: String
		ipAddress: String
		createdAt: Time
	}
	type UsersConnection {
		totalCount: Int!
		edges: [UsersEdge]
		pageInfo: PageInfo!
	}
	type UsersEdge {
		cursor: ID!
		node: User
	}
	type PageInfo {
		startCursor: ID
		endCursor: ID
		hasNextPage: Boolean!
	}
`

type userResolver struct {
	u *model.User
}

//type usersConnectionArgs struct {
//	First *int32
//	After *string
//}

func (r *userResolver) ID() graphql.ID {
	return graphql.ID(strconv.FormatInt(r.u.ID, 10))
}

func (r *userResolver) Email() *string {
	return &r.u.Email
}

func (r *userResolver) Password() *string {
	return &r.u.Password
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
	users []*model.User
	totalCount int
	from  int
	to    int
}

func (r *usersConnectionResolver) TotalCount() int32 {
	return int32(r.totalCount)
}

func (r *usersConnectionResolver) Edges() *[]*usersEdgeResolver {
	l := make([]*usersEdgeResolver, r.to-r.from+1)
	for i := range l {
		l[i] = &usersEdgeResolver{
			cursor: encodeCursor(r.from + i),
			model:  r.users[i],
		}
	}
	return &l
}

func (r *usersConnectionResolver) PageInfo() *pageInfoResolver {
	return &pageInfoResolver{
		startCursor: encodeCursor(r.from),
		endCursor:   encodeCursor(r.to),
		hasNextPage: r.to < r.totalCount,
	}
}

func encodeCursor(i int) graphql.ID {
	return graphql.ID(base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("cursor%d", i+1))))
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
