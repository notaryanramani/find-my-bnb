package store

import (
	"database/sql"
	"log"
)

type Store struct {
	DB    *sql.DB
	User  *UserStore
	Room  *RoomStore
	Hosts *HostStore
}

func NewStore() *Store {
	db, err := NewDB()
	if err != nil {
		log.Fatal(err)
	}

	return &Store{
		DB: db,
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
