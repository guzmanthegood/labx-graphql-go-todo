package resolver

import (
	"labx-graphql-go-todo/model"

	"github.com/guzmanweb/graphql-go"
)

type TodoResolver struct{
	t model.Todo
}

func (r *TodoResolver) ID() graphql.ID {
	return graphql.ID(r.t.ID)
}

func (r *TodoResolver) Text() string {
	return r.t.Text
}

func (r *TodoResolver) User() *UserResolver {
	return &UserResolver{}
}