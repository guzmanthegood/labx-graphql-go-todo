package model

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // pg driver
)

// DataStore interface
type DataStore interface {
	GetUser(id int32) (*User, error)
	AllUsers() ([]User, error)
	CreateUser(name string) (*User, error)
	UpdateUser(id int32, name string) (*User, error)
	DeleteUser(id int32) error
	GetTodo(id int32) (*Todo, error)
	CreateTodo(text string, userID int32) (*Todo, error)
	UpdateTodo(id int32, text, status string) (*Todo, error)
	DeleteTodo(id int32) error
}

// Database struct
type store struct {
	db *sqlx.DB // read database
}

// NewDB : new sqlx client constructor
func NewDataStore(connString string) (DataStore, error) {
	db, err := sqlx.Connect("postgres", connString)
	if err != nil {
		return nil, err
	}
	return store{db}, nil
}
