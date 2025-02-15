package main

import (
	"fmt"
	"github.com/notaryanramani/find-my-bnb/api/vectordb"
)

func main(){
	fmt.Println("Started..")
	db := vectordb.InitVectorDB(300)
	fmt.Println("DB Initialized..")

	db.AddNode("Hello")
	fmt.Printf("Size of DB: %d\n", db.Size())
}