package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	_ "github.com/lib/pq"
	"github.com/notaryanramani/find-my-bnb/api/store"
	"github.com/notaryanramani/find-my-bnb/api/utils"
	"github.com/notaryanramani/find-my-bnb/api/vectordb"
)

type Server struct {
	router   *chi.Mux
	port     string
	store    *store.Store
	cors     *cors.Cors
	vectordb *vectordb.VectorDB
}

func NewServer(port string) *Server {
	return &Server{
		router:   chi.NewRouter(),
		port:     port,
		store:    store.NewStore(),
		cors:     utils.GetCorsMiddleware(),
		vectordb: vectordb.NewVectorDB(),
	}
}

func (s *Server) Run() {
	// Initalize VectorDB
	s.vectordb.InitVectorDB(s.store.DB)

	// Get Routes
	s.router.Get("/hello", HelloWord)
	s.router.Get("/check", Check)

	// Post User Routes
	s.router.Post("/register", s.createUserHandler)
	s.router.Post("/login", s.userLoginHandler)

	// Post Room Routes
	s.router.Post("/rooms", AuthMiddleware(s.getRandomRoomsHandler))

	// Protected Routes
	s.router.Get("/protected", AuthMiddleware(s.protectedHandler))
}
