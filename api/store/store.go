package store

import (
	"log"
)

type Store struct{
	UserStore *UserStore
}

func NewStore() *Store {
	db, err := NewDB()
	if err != nil {
		log.Fatal(err)
	}

	return &Store{
		UserStore: &UserStore{
			db: db,
		},
	}
}