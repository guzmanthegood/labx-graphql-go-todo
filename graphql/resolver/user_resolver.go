package resolver

import (
	"labx-graphql-go-todo/model"
	"strconv"

	"github.com/guzmanweb/graphql-go"
)

type UserResolver struct{
	u model.User
}

func (r *UserResolver) ID() graphql.ID {
	return graphql.ID(strconv.Itoa(int(r.u.ID)))
}

func (r *UserResolver) Name() string {
	return r.u.Name
}

func (r *UserResolver) Todos() []*TodoResolver {
	return nil
}
