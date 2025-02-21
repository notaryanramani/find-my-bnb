package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := ":8080"
	server := NewServer(port)
	server.Run()

	go http.ListenAndServe(server.port, server.cors.Handler(server.router))
	fmt.Println("Server is running on port: ", server.port)
	select {}
}
