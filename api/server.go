package main

import (
	"github.com/go-chi/chi/v5"
)

type Server struct {
	router *chi.Mux
	port   string
}

func NewServer(port string) *Server {
	return &Server{
		router: chi.NewRouter(),
		port:   port,
	}
}

func (s *Server) Init() {
	s.router.Get("/hello", HelloWord)
	s.router.Post("/register", validateUserCredentials(RegisterUser))
}
