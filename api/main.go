package main

import (
	"fmt"
	"net/http"
)

func main(){
	port := ":8080"
	server := NewServer(port)
	server.Run()

	
	
	go http.ListenAndServe(server.port, server.cors.Handler(server.router))
	fmt.Println("Server is running on port: ", server.port)
	select {}
}