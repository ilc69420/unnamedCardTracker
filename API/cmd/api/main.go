package main

import (
	"cardforge/internal/server"
	"log"
	"log/slog"
)

func main() {

	srv := server.New()

	slog.Info("server started")
	if err := srv.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
