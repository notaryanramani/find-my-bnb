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

func (u *UserStore) AddRoomToWishlist(ctx context.Context, username string, roomID int64) error {
	query := `UPDATE users SET wishlist = array_append(wishlist, $1) WHERE username = $2`

	_, err := u.db.ExecContext(ctx, query, roomID, username)

	return err
}

func (u *UserStore) RemoveRoomFromWishlist(ctx context.Context, username string, roomID int64) error {
	query := `UPDATE users SET wishlist = array_remove(wishlist, $1) WHERE username = $2`

	_, err := u.db.ExecContext(ctx, query, roomID, username)

	return err
}

func (u *UserStore) GetWishlist(ctx context.Context, username string) ([]int64, error) {
	query := `SELECT wishlist FROM users WHERE username = $1`

	var wishlist []int64
	err := u.db.QueryRowContext(ctx, query, username).Scan(&wishlist)
	if err != nil {
		return nil, err
	}

	return wishlist, nil
}

func (u *UserStore) CheckRoomExistInWishlist(ctx context.Context, username string, roomID int64) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE username = $1 AND $2 = ANY(wishlist));`

	var exist bool
	err := u.db.QueryRowContext(ctx, query, username, roomID).Scan(&exist)

	return exist, err
}
