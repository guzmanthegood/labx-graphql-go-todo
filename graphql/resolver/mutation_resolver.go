package resolver

import (
	"context"
	"github.com/guzmanweb/graphql-go"
	"labx-graphql-go-todo/model"
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

func (r *TodoMutationResolver) CreateUser(ctx context.Context, args *struct {
	Input *CreateUserInput
}) *UserResolver {
	return &UserResolver{u:model.User{
		ID:   "u14",
		Name: args.Input.Name,
	}}
}

func (r *TodoMutationResolver) UpdateUser(ctx context.Context, args *struct {
	Input *UpdateUserInput
}) *UserResolver {
	return &UserResolver{u:model.User{
		ID:   string(args.Input.ID),
		Name: args.Input.Name,
	}}
}

func (r *TodoMutationResolver) DeleteUser(ctx context.Context, args *struct {
	Input *DeleteUserInput
}) *UserResolver {
	return &UserResolver{u:model.User{}}
}
