package app

import (
	"database/sql"
)

type apiserver struct {
	db       *sql.DB
	Sessions map[string]*Session
}
