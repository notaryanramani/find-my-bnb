package store

import (
	"log"
)

type Store struct{
	User *UserStore
	Room *RoomStore
	Hosts *HostStore
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
		Room: &RoomStore{
			db: db,
		},
		Hosts: &HostStore{
			db: db,
		},
	}
}