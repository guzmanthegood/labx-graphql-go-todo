package resolver

import (
	"context"
	"fmt"
	"strconv"

	"labx-graphql-go-todo/model"

	"github.com/guzmanweb/graphql-go"
)

type UserFilter struct {
	ID graphql.ID
}

type TodoFilter struct {
	Code graphql.ID
}


func (r *TodoQueryResolver) User(ctx context.Context, args *UserFilter) *UserResolver {

	ID, err := strconv.Atoi(string(args.ID))
	if err != nil {
		panic(err)
	}

	user, err := model.GetDataStore().GetUser(int32(ID))
	if err != nil {
		panic(err)
	}
	return &UserResolver{u: user}
}

func (r *TodoQueryResolver) AllUsers(ctx context.Context) []*UserResolver {

	users, err := model.GetDataStore().AllUsers()
	if err != nil {
		panic(err)
	}
	var usersResolvers []*UserResolver
	for _, u := range users {
		usersResolvers = append(usersResolvers, &UserResolver{u	})
	}
	return usersResolvers
}

func (r *TodoQueryResolver) Todo(ctx context.Context, args *TodoFilter) *TodoResolver {
	return &TodoResolver{model.Todo{
		ID:   string(args.Code),
		Text: fmt.Sprintf("Text %v", args.Code),
	}}
}

func (r *TodoQueryResolver) AllTodos(ctx context.Context) []*TodoResolver {
	return nil
}