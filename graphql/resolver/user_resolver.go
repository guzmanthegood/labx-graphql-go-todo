package resolver

import (
	"github.com/guzmanweb/graphql-go"
	"labx-graphql-go-todo/model"
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
