package app

import (
	"cardforge/internal/database"
	"cardforge/pkg/logger"
	"database/sql"
)

type App struct {
	DB     *sql.DB
	Logger logger.Logger
}

func New() *App {
	return &App{
		DB:     database.New(),
		Logger: logger.New(),
	}
}
