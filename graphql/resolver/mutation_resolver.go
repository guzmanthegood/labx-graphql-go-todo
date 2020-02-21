package resolver

import (
	"context"
	"github.com/guzmanweb/graphql-go"
	"labx-graphql-go-todo/model"
	"log"
	"strconv"
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
	user, err := model.GetDataStore().CreateUser(args.Input.Name)
	if err != nil{
		panic(err)
	}
	if user != nil {
		return &UserResolver{model.User{
			ID:   user.ID,
			Name: user.Name,
		}}
	}
	return &UserResolver{}
}

func (r *TodoMutationResolver) UpdateUser(ctx context.Context, args *struct {
	Input *UpdateUserInput
}) *UserResolver {
	ID, err := strconv.Atoi(string(args.Input.ID))
	if err != nil {
		panic(err)
	}

	user, err := model.GetDataStore().UpdateUser(int32(ID), args.Input.Name)
	if err != nil{
		panic(err)
	}
	if user != nil{
		return &UserResolver{model.User{
			ID:   user.ID,
			Name: user.Name,
		}}
	}
	return &UserResolver{}
}

func (r *TodoMutationResolver) DeleteUser(ctx context.Context, args *struct {
	Input *DeleteUserInput
}) *UserResolver {
	ID, err := strconv.Atoi(string(args.Input.ID))
	if err != nil {
		panic(err)
	}
	err = model.GetDataStore().DeleteUser(int32(ID))
	if err !=  nil {
		panic(err)
	}
	return &UserResolver{}
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