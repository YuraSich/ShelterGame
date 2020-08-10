package app

import (
	"database/sql"

	_ "github.com/lib/pq" //
)

// Apiserver ...
type Apiserver struct {
	db       *sql.DB
	sessions map[string]*Session
}

// NewAPIServer ...
func NewAPIServer(dbName string) (*Apiserver, error) {
	a := &Apiserver{
		db: nil,
	}
	if err := a.connectToDB(dbName); err != nil {
		return nil, err
	}
	return a, nil
}

// connectToDB ....
func (a *Apiserver) connectToDB(dbName string) error {
	db, err := sql.Open("postgres", "host=localhost dbname="+dbName+" sslmode=disable")
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	a.db = db
	return nil
}
