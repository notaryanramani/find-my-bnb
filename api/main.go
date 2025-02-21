package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := chi.NewRouter()

	port := ":8080"
	server := NewServer(port)
	server.Run()

	r.Mount("/api", server.router)

	go http.ListenAndServe(server.port, server.cors.Handler(r))
	fmt.Println("Server is running on port: ", server.port)
	select {}
}
