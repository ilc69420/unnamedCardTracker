package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func New() *sql.DB {
	db, err := sql.Open("sqlite3", "../cardForge.db")
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	return db
}
