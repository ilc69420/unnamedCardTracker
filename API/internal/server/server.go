package server

import (
	"cardforge/internal/app"
	"log"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func New() *Server {

	app := app.New()

	s := &http.Server{
		Handler: registerRoutes(app),
		Addr:    ":4000",
	}

	return &Server{httpServer: s}
}

func (s *Server) Run() error {
	log.Println("Server listening on", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}
