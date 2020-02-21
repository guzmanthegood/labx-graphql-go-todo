package model

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // pg driver
)

// DataStore interface
type DataStore interface {
	CreateUser(name string) (*User, error)
	UpdateUser(id int32, name string) (*User, error)
	DeleteUser(id int32) error
}

// Database struct
type store struct {
	db *sqlx.DB // read database
}

// NewDB : new sqlx client constructor
// write/read from the same source
func NewDataStore(connString string) (DataStore, error) {
	db, err := sqlx.Connect("postgres", connString)
	if err != nil {
		return nil, err
	}
	return store{db}, nil
}
