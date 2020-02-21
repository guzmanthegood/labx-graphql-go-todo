package resolver

import (
	"labx-graphql-go-todo/model"
	"strconv"

	"github.com/guzmanweb/graphql-go"
)

type TodoResolver struct{
	t model.Todo
}

func (r *TodoResolver) ID() graphql.ID {
	return graphql.ID(strconv.Itoa(int(r.t.ID)))
}

func (r *TodoResolver) Text() string {
	return r.t.Text
}

func (r *TodoResolver) Status() string {
	return r.t.Status
}

func (r *TodoResolver) User() *UserResolver {
	user, err := model.GetDataStore().GetUser(r.t.UserID)
	if err != nil {
		panic(err)
	}
	if user == nil {
		return &UserResolver{}
	}
	return &UserResolver{*user}
}