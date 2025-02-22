package store

import (
	"fmt"
)

func CreateRoomPayloadFromRoomResponse(room *Room) *RoomPayload {
	var description string
	if room.Description.Valid {
		description = room.Description.String
	} else {
		description = "N/A"
	}

	var neighborhoodOverview string
	if room.NeighborhoodOverview.Valid {
		neighborhoodOverview = room.NeighborhoodOverview.String
	} else {
		neighborhoodOverview = "N/A"
	}

	var price float32
	if room.Price.Valid {
		price = float32(room.Price.Float64)
	} else {
		price = -1
	}

	var bedrooms int
	if room.Bedrooms.Valid {
		bedrooms = int(room.Bedrooms.Int64)
	} else {
		bedrooms = -1
	}

	var beds int
	if room.Beds.Valid {
		beds = int(room.Beds.Int64)
	} else {
		beds = -1
	}

	var neighborhood string
	if room.Neighborhood.Valid {
		neighborhood = room.Neighborhood.String
	} else {
		neighborhood = "N/A"
	}

	IDString := fmt.Sprintf("%d", room.ID)

	return &RoomPayload{
		ID:                   room.ID,
		IDString:             IDString,
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
