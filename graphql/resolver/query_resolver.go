package resolver

import (
	"context"
	"fmt"
	"log"

	"labx-graphql-go-todo/model"

	"github.com/guzmanweb/graphql-go"
)

type UserFilter struct {
	Code graphql.ID
}

type TodoFilter struct {
	Code graphql.ID
}


func (r *TodoQueryResolver) User(ctx context.Context, args *UserFilter) *UserResolver {
	log.Printf("[INFO] query/todo/user(code:%v)", args.Code)
	return &UserResolver{u:model.User{
		ID:   string(args.Code),
		Name: fmt.Sprintf("Perico %v", args.Code),
	}}
}

func (r *TodoQueryResolver) AllUsers(ctx context.Context) []*UserResolver {
	log.Println("[INFO] query/todo/allUsers")
	return nil
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