package main

import (
	"database/sql"

	"github.com/YuraSich/ShelterGame/app"
)

type apiserver struct {
	db       *sql.DB
	Sessions map[string]*app.Session
}


