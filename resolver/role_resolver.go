package resolver

import (
	"github.com/OscarYuen/go-graphql-starter/model"
	"github.com/neelance/graphql-go"
	"strconv"
)

type roleResolver struct {
	role *model.Role
}

func (r *roleResolver) ID() graphql.ID {
	return graphql.ID(strconv.FormatInt(r.role.ID, 10))
}

func (r *roleResolver) Name() *string {
	return &r.role.Name
}
