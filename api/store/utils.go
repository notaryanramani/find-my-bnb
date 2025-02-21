package store

// type RoomPayload struct {
// 	ID                   int64   `json:"id"`
// 	ListingURL           string  `json:"listing_url"`
// 	Name                 string  `json:"name"`
// 	Description          string  `json:"description"`
// 	NeighborhoodOverview string  `json:"neighborhood_overview"`
// 	PictureURL           string  `json:"picture_url"`
// 	Price                float32 `json:"price"`
// 	Bedrooms             int     `json:"bedrooms"`
// 	Beds                 int     `json:"beds"`
// 	RoomType             string  `json:"room_type"`
// 	PropertyType         string  `json:"property_type"`
// 	Neighborhood         string  `json:"neighborhood"`
// 	HostID               int64   `json:"host_id"`
// }

func CreateRoomPayloadFromRoomResponse(room *Room) *RoomPayload {
	var description string
	if room.Description.Valid {
		description = room.Description.String
	} else {
		description = ""
	}

	var neighborhoodOverview string
	if room.NeighborhoodOverview.Valid {
		neighborhoodOverview = room.NeighborhoodOverview.String
	} else {
		neighborhoodOverview = ""
	}

	var price float32
	if room.Price.Valid {
		price = float32(room.Price.Float64)
	} else {
		price = 0
	}

	var bedrooms int
	if room.Bedrooms.Valid {
		bedrooms = int(room.Bedrooms.Int64)
	} else {
		bedrooms = 0
	}

	var beds int
	if room.Beds.Valid {
		beds = int(room.Beds.Int64)
	} else {
		beds = 0
	}

	var neighborhood string
	if room.Neighborhood.Valid {
		neighborhood = room.Neighborhood.String
	} else {
		neighborhood = ""
	}

	return &RoomPayload{
		ID:                   room.ID,
		ListingURL:           room.ListingURL,
		Name:                 room.Name,
		Description:          description,
		NeighborhoodOverview: neighborhoodOverview,
		PictureURL:           room.PictureURL,
		Price:                price,
		Bedrooms:             bedrooms,
		Beds:                 beds,
		RoomType:             room.RoomType,
		PropertyType:         room.PropertyType,
		Neighborhood:         neighborhood,
		HostID:               room.HostID,
	}
}