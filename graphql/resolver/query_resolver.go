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