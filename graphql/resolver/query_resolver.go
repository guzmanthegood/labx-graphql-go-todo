package resolver

import (
	"context"
	"fmt"
	"log"
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
	log.Printf("[INFO] query/todo/user(id:%v)", args.ID)

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
	log.Println("[INFO] query/todo/allUsers")

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
	log.Printf("[INFO] query/todo/todo(code:%v)", args.Code)
	return &TodoResolver{model.Todo{
		ID:   string(args.Code),
		Text: fmt.Sprintf("Text %v", args.Code),
	}}
}

func (r *TodoQueryResolver) AllTodos(ctx context.Context) []*TodoResolver {
	log.Println("[INFO] query/todo/allTodos")
	return nil
}