package store

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func NewDB() (*sql.DB, error) {
	addr, ok := os.LookupEnv("DB_ADDR")
	if !ok {
		log.Fatal("DB_ADDR is not set in .env file")
		return nil, fmt.Errorf("DB_ADDR is not set in .env file")
	}
	db, err := sql.Open("postgres", addr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	idleTime, err := time.ParseDuration("15m")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	db.SetConnMaxIdleTime(idleTime)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return db, nil	
}
