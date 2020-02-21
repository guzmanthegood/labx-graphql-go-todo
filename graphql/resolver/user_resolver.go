package resolver

import (
	"labx-graphql-go-todo/model"

	"github.com/guzmanweb/graphql-go"
)

type UserResolver struct{
	u model.User
}

func (r *UserResolver) ID() graphql.ID {
	return graphql.ID(r.u.ID)
}

func (r *UserResolver) Name() string {
	return r.u.Name
}
