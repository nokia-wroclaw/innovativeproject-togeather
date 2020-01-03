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
		Location *Location `json:"location"`
	}

	Location struct {
		GeoLat float64 `json:"lat" db:"geolat"`
		GeoLon float64 `json:"lon" db:"geolon"`
		Address string `json:"lobby_address"`
	}


	LobbyService interface {
		List(ctx context.Context) ([]*Lobby, error)

		Create(
			ctx context.Context,
			restaurantID int,
			ownerID int,
			expires *time.Time,
			address string,
		) (*Lobby, error)
	}

	LobbyStore interface {
		List(ctx context.Context) ([]*Lobby, error)

		Create(
			ctx context.Context,
			restaurantID int,
			ownerID int,
			expires *time.Time,
			address string,
		) (*Lobby, error)
	}
)
