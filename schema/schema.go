package schema

import (
	graphql "github.com/neelance/graphql-go"
	"log"
	"strconv"
)

var Schema = `
	schema {
		query: Query
		mutation: Mutation
	}
	type Query {
		user(id: ID!): User
	}
	type Mutation {
		createUser(name: String!, password: String!): User
	}
	type User {
		id: ID!
		name: String!
		password: String!
	}`

type user struct {
	ID        graphql.ID
	Name      string
	Password  string
}

var users = []*user{
	{
		ID:        "1",
		Name:      "Luke Skywalker",
		Password:  "1234",
	},
}
var userData = make(map[graphql.ID]*user)

func init() {
	for _, user := range users {
		userData[user.ID] = user
	}
}

type Resolver struct{}

type userResolver struct {
	u *user
}

func (r *userResolver) ID() graphql.ID {
	return r.u.ID
}

func (r *userResolver) Name() string {
	return r.u.Name
}

func (r *userResolver) Password() string{
	return r.u.Password
}


func (r *Resolver) User(args struct{ ID graphql.ID }) *userResolver {
	if user := userData[args.ID]; user != nil {
		return &userResolver{user}
	}
	return nil
}

func (r *Resolver) CreateUser(args *struct {
	Name string
	Password  string
}) *userResolver {
	log.Println( len(users) + 1)
	user := &user{
		ID: graphql.ID(strconv.Itoa(len(users)+1)),
		Name:args.Name,
		Password:args.Password,
	}
	users = append(users, user)
	userData[user.ID] = user
	return &userResolver{user}
}