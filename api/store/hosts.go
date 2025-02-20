package store

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

type HostStore struct {
	db *sql.DB
}

type Host struct {
	ID              int
	HostURL         string
	HostName        string
	HostSince       time.Time
	HostLocation    string
	HostAbout       string
	HostThumnailURL string
	HostPictureURL  string
}

type HostPayload struct {
	ID              int       `json:"id"`
	HostURL         string    `json:"host_url"`
	HostName        string    `json:"host_name"`
	HostSince       time.Time `json:"host_since"`
	HostLocation    string    `json:"host_location"`
	HostAbout       string    `json:"host_about"`
	HostThumnailURL string    `json:"host_thumnail_url"`
	HostPictureURL  string    `json:"host_picture_url"`
}

func (h *HostStore) Create(ctx context.Context, host *Host) error {
	query := `INSERT INTO 
	hosts (id, host_url, host_name, host_since, host_location, host_about, host_thumnail_url, host_picture_url) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	err := h.db.QueryRowContext(
		ctx,
		query,
		host.ID,
		host.HostURL,
		host.HostName,
		host.HostSince,
		host.HostLocation,
		host.HostAbout,
		host.HostThumnailURL,
		host.HostPictureURL,
	).Scan()

	return err
}

func (h *HostStore) GetByID(ctx context.Context, id int) (*Host, error) {
	query := `SELECT * FROM hosts WHERE id = $1`

	host := &Host{}
	err := h.db.QueryRowContext(ctx, query, id).Scan(
		&host.ID,
		&host.HostURL,
		&host.HostName,
		&host.HostSince,
		&host.HostLocation,
		&host.HostAbout,
		&host.HostThumnailURL,
		&host.HostPictureURL,
	)

	return host, err
}

func (h *HostStore) Update(ctx context.Context, host *Host) error {
	query := `UPDATE hosts SET host_url = $1, host_name = $2, host_since = $3, host_location = $4, host_about = $5, host_thumnail_url = $6, host_picture_url = $7 WHERE id = $8`

	_, err := h.db.ExecContext(
		ctx,
		query,
		host.HostURL,
		host.HostName,
		host.HostSince,
		host.HostLocation,
		host.HostAbout,
		host.HostThumnailURL,
		host.HostPictureURL,
		host.ID,
	)

	return err
}

func (h *HostStore) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM hosts WHERE id = $1`

	_, err := h.db.ExecContext(ctx, query, id)

	return err
}
