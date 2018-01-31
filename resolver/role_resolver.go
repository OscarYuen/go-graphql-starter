package resolver

import (
	"github.com/OscarYuen/go-graphql-starter/model"
	"github.com/neelance/graphql-go"
)

type roleResolver struct {
	role *model.Role
}

func (r *roleResolver) ID() graphql.ID {
	return graphql.ID(r.role.ID)
}

func (r *roleResolver) Name() *string {
	return &r.role.Name
}
