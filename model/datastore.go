package model

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // pg driver
)

// DataStore interface
type DataStore interface {}

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
