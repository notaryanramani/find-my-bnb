package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	_ "github.com/lib/pq"
	"github.com/notaryanramani/find-my-bnb/api/store"
	"github.com/notaryanramani/find-my-bnb/api/utils"
)

type Server struct {
	router *chi.Mux
	port   string
	store  *store.Store
	cors *cors.Cors
}

func NewServer(port string) *Server {
	return &Server{
		router: chi.NewRouter(),
		port:   port,
		store:  store.NewStore(),
		cors: utils.GetCorsMiddleware(),
	}
}

func (s *Server) Run() {
	// Get Routes
	s.router.Get("/hello", HelloWord)
	s.router.Get("/check", Check)

	// Post Routes
	s.router.Post("/register", s.createUserHandler)
	s.router.Post("/login", s.userLoginHandler)

	// Protected Routes
	s.router.Get("/protected", AuthMiddleware(s.protectedHandler))
}
