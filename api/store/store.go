package store

import (
	"log"
)

type Store struct{
	User *UserStore
}

func NewStore() *Store {
	db, err := NewDB()
	if err != nil {
		log.Fatal(err)
	}

	return &Store{
		User: &UserStore{
			db: db,
		},
	}
}