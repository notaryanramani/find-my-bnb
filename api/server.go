package main

import (
	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
	"github.com/notaryanramani/find-my-bnb/api/store"
)

type Server struct {
	router *chi.Mux
	port   string
	store  *store.Store
}

func NewServer(port string) *Server {
	return &Server{
		router: chi.NewRouter(),
		port:   port,
		store:  store.NewStore(),
	}
}

func (s *Server) Run() {
	// Get Routes
	s.router.Get("/hello", HelloWord)
	s.router.Get("/check", Check)

	// Post Routes
	s.router.Post("/register", s.createUserHandler)

	// Protected Routes
	s.router.Post("/login", AuthMiddleware(s.userLoginHandler))
}
