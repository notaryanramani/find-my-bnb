package store

import (
	"context"
	"database/sql"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

type RoomStore struct {
	db *sql.DB
}

type Room struct {
	ID                   int64
	ListingURL           string
	Name                 string
	Description          sql.NullString
	NeighborhoodOverview sql.NullString
	PictureURL           string
	Price                sql.NullFloat64
	Bedrooms             sql.NullInt64
	Beds                 sql.NullInt64
	RoomType             string
	PropertyType         string
	Neighborhood         sql.NullString
	HostID               int64
}

type RoomPayload struct {
	ID                   int64   `json:"id"`
	IDString             string  `json:"id_string"`
	ListingURL           string  `json:"listing_url"`
	Name                 string  `json:"name"`
	Description          string  `json:"description"`
	NeighborhoodOverview string  `json:"neighborhood_overview"`
	PictureURL           string  `json:"picture_url"`
	Price                float32 `json:"price"`
	Bedrooms             int     `json:"bedrooms"`
	Beds                 int     `json:"beds"`
	RoomType             string  `json:"room_type"`
	PropertyType         string  `json:"property_type"`
	Neighborhood         string  `json:"neighborhood"`
	HostID               int64   `json:"host_id"`
}

type TopKPayload struct {
	K   int     `json:"k"`
	Ids []int64 `json:"ids"`
}

type RoomsPayload struct {
	Rooms []*RoomPayload `json:"rooms"`
}

func (r *RoomStore) Create(ctx context.Context, room *Room) error {
	query := `INSERT INTO rooms (id, listing_url, name, description, neighborhood_overview, picture_url, price, bedrooms, beds, room_type, property_type, neighborhood, host_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`

	err := r.db.QueryRowContext(
		ctx,
		query,
		room.ID,
		room.ListingURL,
		room.Name,
		room.Description,
		room.NeighborhoodOverview,
		room.PictureURL,
		room.Price,
		room.Bedrooms,
		room.Beds,
		room.RoomType,
		room.PropertyType,
		room.Neighborhood,
		room.HostID,
	).Scan()

	return err
}

func (r *RoomStore) GetTopKRandom(ctx context.Context, k int) ([]*RoomPayload, error) {
	query := `SELECT * from ROOMS ORDER BY RANDOM() LIMIT $1`

	rows, err := r.db.QueryContext(ctx, query, k)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rooms := []*RoomPayload{}
	for rows.Next() {
		room := &Room{}
		err := rows.Scan(
			&room.ID,
			&room.ListingURL,
			&room.Name,
			&room.Description,
			&room.NeighborhoodOverview,
			&room.PictureURL,
			&room.Price,
			&room.Bedrooms,
			&room.Beds,
			&room.RoomType,
			&room.PropertyType,
			&room.Neighborhood,
			&room.HostID,
		)
		if err != nil {
			return nil, err
		}

		roomPayload := CreateRoomPayloadFromRoomResponse(room)
		rooms = append(rooms, roomPayload)
	}
	return rooms, nil
}

func (r *RoomStore) NextTopKRandom(ctx context.Context, k int, ids []int64) ([]*RoomPayload, error) {
	// Create placeholder for the query
	placeholder := make([]string, len(ids))
	for i := range ids {
		placeholder[i] = "$" + strconv.Itoa(i+2)
	}

	query := `SELECT * from ROOMS WHERE id NOT IN (` + strings.Join(placeholder, ",") + `) ORDER BY RANDOM() LIMIT $1`

	// Create interface for K value and IDS value
	args := make([]interface{}, len(ids)+1)
	args[0] = k
	for i, id := range ids {
		args[i+1] = id
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rooms := []*RoomPayload{}
	for rows.Next() {
		room := &Room{}
		err := rows.Scan(
			&room.ID,
			&room.ListingURL,
			&room.Name,
			&room.Description,
			&room.NeighborhoodOverview,
			&room.PictureURL,
			&room.Price,
			&room.Bedrooms,
			&room.Beds,
			&room.RoomType,
			&room.PropertyType,
			&room.Neighborhood,
			&room.HostID,
		)
		if err != nil {
			return nil, err
		}
		roomPayload := CreateRoomPayloadFromRoomResponse(room)
		rooms = append(rooms, roomPayload)
	}
	return rooms, nil
}

func (r *RoomStore) GetByID(ctx context.Context, id int64) (*RoomPayload, error) {
	query := `SELECT * FROM rooms WHERE id = $1`

	room := &Room{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&room.ID,
		&room.ListingURL,
		&room.Name,
		&room.Description,
		&room.NeighborhoodOverview,
		&room.PictureURL,
		&room.Price,
		&room.Bedrooms,
		&room.Beds,
		&room.RoomType,
		&room.PropertyType,
		&room.Neighborhood,
		&room.HostID,
	)
	if err != nil {
		return nil, err
	}

	roomPayload := CreateRoomPayloadFromRoomResponse(room)

	return roomPayload, nil
}

func (r *RoomStore) GetByMultipleIDs(ctx context.Context, ids []int64) ([]*RoomPayload, error) {
	// Create placeholder for the query
	placeholder := make([]string, len(ids))
	for i := range ids {
		placeholder[i] = "$" + strconv.Itoa(i+1)
	}

	query := `SELECT * FROM rooms WHERE id IN (` + strings.Join(placeholder, ",") + `)`

	// Create interface for IDS value
	args := make([]interface{}, len(ids))
	for i, id := range ids {
		args[i] = id
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rooms := []*RoomPayload{}
	for rows.Next() {
		room := &Room{}
		err := rows.Scan(
			&room.ID,
			&room.ListingURL,
			&room.Name,
			&room.Description,
			&room.NeighborhoodOverview,
			&room.PictureURL,
			&room.Price,
			&room.Bedrooms,
			&room.Beds,
			&room.RoomType,
			&room.PropertyType,
			&room.Neighborhood,
			&room.HostID,
		)
		if err != nil {
			return nil, err
		}

		roomPayload := CreateRoomPayloadFromRoomResponse(room)
		rooms = append(rooms, roomPayload)
	}
	return rooms, nil
}

func (r *RoomStore) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM rooms WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *RoomStore) Update(ctx context.Context, room *Room) error {
	query := `UPDATE rooms SET 
	listing_url = $1, 
	name = $2, 
	description = $3, 
	neighborhood_overview = $4, 
	picture_url = $5, 
	price = $6, 
	bedrooms = $7, 
	beds = $8,
	room_type = $9,
	property_type = $10,
	neighborhood = $11,
	host_id = $12
	WHERE id = $13`

	_, err := r.db.ExecContext(
		ctx,
		query,
		room.ListingURL,
		room.Name,
		room.Description,
		room.NeighborhoodOverview,
		room.PictureURL,
		room.Price,
		room.Bedrooms,
		room.Beds,
		room.RoomType,
		room.PropertyType,
		room.Neighborhood,
		room.HostID,
		room.ID,
	)
	return err
}
