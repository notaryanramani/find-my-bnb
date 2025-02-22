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
	Username string `json:"username"`
	Message  string `json:"message"`
}

type UserLoginPayload struct {
	Username string `json:"username"`
	Token    string `json:"token"`
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

func (u *UserStore) Update(ctx context.Context, user *User) error {
	query := `UPDATE users SET name = $1, username = $2, email = $3, password = $4 WHERE id = $5`

	_, err := u.db.ExecContext(
		ctx,
		query,
		user.Name,
		user.Username,
		user.Email,
		user.Password,
		user.ID,
	)

	return err
}

func (u *UserStore) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM users WHERE id = $1`

	_, err := u.db.ExecContext(ctx, query, id)

	return err
}
