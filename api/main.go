package main

import (
	"fmt"
	"net/http"
)

func main(){
	port := ":8080"
	server := NewServer(port)
	server.Init()
	go http.ListenAndServe(server.port, server.router)
	fmt.Println("Server is running on port: ", server.port)
	select {}
}