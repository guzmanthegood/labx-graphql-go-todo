package resolver

import (
	"context"
	"errors"
	"strconv"

	"labx-graphql-go-todo/model"

	"github.com/guzmanweb/graphql-go"
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
	Status	*string
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
	userID, err := strconv.Atoi(string(args.Input.UserID))
	if err != nil {
		panic(err)
	}
	user, err := model.GetDataStore().GetUser(int32(userID))
	if err != nil {
		panic(err)
	}
	if user == nil {
		panic(errors.New("user not found"))
	}

	todo, err := model.GetDataStore().CreateTodo(args.Input.Text, user.ID)
	if err != nil{
		panic(err)
	}
	if todo != nil {
		return &TodoResolver{*todo}
	}
	return &TodoResolver{}
}

func (r *TodoMutationResolver) UpdateTodo(ctx context.Context, args *struct {
	Input *UpdateTodoInput
}) *TodoResolver {
	todoID, err := strconv.Atoi(string(args.Input.ID))
	if err != nil {
		panic(err)
	}

	todo, err := model.GetDataStore().GetTodo(int32(todoID))
	if err != nil{
		panic(err)
	}
	if todo == nil {
		panic(errors.New("todo not found"))
	}

	text := todo.Text
	status := todo.Status
	if args.Input.Text != nil {
		text = *args.Input.Text
	}
	if args.Input.Status != nil {
		status = *args.Input.Status
	}

	todo, err = model.GetDataStore().UpdateTodo(int32(todoID), text, status)
	if err != nil{
		panic(err)
	}
	if todo != nil{
		return &TodoResolver{*todo}
	}
	return &TodoResolver{}
}

func (r *TodoMutationResolver) DeleteTodo(ctx context.Context, args *struct {
	Input *DeleteTodoInput
}) *TodoResolver {
	todoID, err := strconv.Atoi(string(args.Input.ID))
	if err != nil {
		panic(err)
	}

	todo, err := model.GetDataStore().GetTodo(int32(todoID))
	if err != nil{
		panic(err)
	}
	if todo == nil {
		panic(errors.New("todo not found"))
	}

	err = model.GetDataStore().DeleteTodo(int32(todoID))
	if err !=  nil {
		panic(err)
	}
	return &TodoResolver{}
}