package store

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

type UserStore struct {
	db *sql.DB
}

type User struct {
	ID       int
	Name     string
	Username string
	Email    string
	Password []byte
}

type UserJSON struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserPayload struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type UserLoginPayload struct {
	Username string `json:"username"`
	Token	string `json:"token"`
}

func (u *UserStore) Create(ctx context.Context, user *User) error {
	query := `INSERT INTO users (name, username, email, password) VALUES ($1, $2, $3, $4) RETURNING id`

	err := u.db.QueryRowContext(
		ctx,
		query,
		user.Name,
		user.Username,
		user.Email,
		user.Password,
	).Scan(&user.ID)

	return err
}

func (u *UserStore) GetByUsername(ctx context.Context, username string) (*User, error) {
	query := `SELECT id, name, username, email, password FROM users WHERE username = $1`

	user := &User{}
	err := u.db.QueryRowContext(ctx, query, username).Scan(
		&user.ID,
		&user.Name,
		&user.Username,
		&user.Email,
		&user.Password,
	)

	return user, err
}