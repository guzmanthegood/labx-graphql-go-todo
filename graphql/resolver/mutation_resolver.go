package resolver

import (
	"context"
	"github.com/guzmanweb/graphql-go"
	"labx-graphql-go-todo/model"
	"log"
)

type CreateUserInput struct {
	Name string
}

type UpdateUserInput struct {
	ID 		graphql.ID
	Name 	string
}

type DeleteUserInput struct {
	ID graphql.ID
}

type CreateTodoInput struct {
	Text 	string
	UserID 	graphql.ID
}

type UpdateTodoInput struct {
	ID 		graphql.ID
	Text 	*string
}

type DeleteTodoInput struct {
	ID graphql.ID
}

func (r *TodoMutationResolver) CreateUser(ctx context.Context, args *struct {
	Input *CreateUserInput
}) *UserResolver {
	return &UserResolver{model.User{
		ID:   "u14",
		Name: args.Input.Name,
	}}
}

func (r *TodoMutationResolver) UpdateUser(ctx context.Context, args *struct {
	Input *UpdateUserInput
}) *UserResolver {
	return &UserResolver{model.User{
		ID:   string(args.Input.ID),
		Name: args.Input.Name,
	}}
}

func (r *TodoMutationResolver) DeleteUser(ctx context.Context, args *struct {
	Input *DeleteUserInput
}) *UserResolver {
	return &UserResolver{model.User{}}
}

func (r *TodoMutationResolver) CreateTodo(ctx context.Context, args *struct {
	Input *CreateTodoInput
}) *TodoResolver {
	log.Printf("[INFO] mutation/todo/createTodo(%v)", args.Input)
	return &TodoResolver{model.Todo{
		ID:   "u14",
		Text: args.Input.Text,
	}}
}

func (r *TodoMutationResolver) UpdateTodo(ctx context.Context, args *struct {
	Input *UpdateTodoInput
}) *TodoResolver {
	return &TodoResolver{model.Todo{
		ID:   string(args.Input.ID),
		Text: *args.Input.Text,
	}}
}

func (r *TodoMutationResolver) DeleteTodo(ctx context.Context, args *struct {
	Input *DeleteTodoInput
}) *TodoResolver {
	return &TodoResolver{model.Todo{}}
}