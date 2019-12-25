package core

import (
	"context"
	"time"
)

type (
	Lobby struct {
		ID int `json:"id" db:"id"`
		Restaurant *Restaurant `json:"restaurant"`
		Owner int `json:"owner" db:"owner"`
		Expires time.Time `json:"expires" db:"expires"`
		GeoLat float64 `json:"lat" db:"geolat"`
		GeoLon float64 `json:"lon" db:"geolon"`
	}


	LobbyService interface {
		List(ctx context.Context) ([]*Lobby, error)

		Create(
			ctx context.Context,
			restaurantID int,
			ownerID int,
			expires *time.Time,
			geolat float64,
			geolon float64,
		) (*Lobby, error)
	}

	LobbyStore interface {
		List(ctx context.Context) ([]*Lobby, error)

		Create(
			ctx context.Context,
			restaurantID int,
			ownerID int,
			expires *time.Time,
			geolat float64,
			geolon float64,
		) (*Lobby, error)
	}
)
