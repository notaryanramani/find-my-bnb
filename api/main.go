package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/cors"
)

func main(){
	port := ":8080"
	server := NewServer(port)
	server.Run()

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
	})

	go http.ListenAndServe(server.port, corsMiddleware.Handler(server.router))
	fmt.Println("Server is running on port: ", server.port)
	select {}
}